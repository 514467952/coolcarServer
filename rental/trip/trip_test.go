package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	poi "coolcar/rental/trip/client"
	"coolcar/rental/trip/dao"
	sharedauth "coolcar/shared/auth"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	mongotesting "coolcar/shared/testing"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestCreateTrip(t *testing.T) {
	c := sharedauth.ContextWithAccountID(
		context.Background(),
		id.AccountID("account1"),
	)
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger:%v", err)
	}

	pm := &profileManager{}
	cm := &carManager{}
	s := &Service{
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		Logger:         logger,
	}

	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Loaction{
			Latitude:  32.123,
			Longitude: 114.2525,
		},
	}
	pm.iID = "identity1"
	golden := `{"accountID":"account1","carID":"car1","start":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门"},"current":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门"},"status":1,"identity_id":"identity1"}`
	cases := []struct {
		name         string
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:   "normal_create",
			tripID: "5f8132eb00714bf62948905c",
			want:   golden,
		},
		{
			name:       "profile_err",
			tripID:     "5f8132eb00714bf629489055",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			tripID:       "5f8132eb00714bf629489056",
			carVerifyErr: fmt.Errorf("verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err", //开锁失败，但是行程创建成功
			tripID:       "5f8132eb00714bf629489057",
			carUnlockErr: fmt.Errorf("unlock"),
			want:         golden,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			//tripid
			mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))

			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyErr

			res, err := s.CreateTrip(c, req)
			if cc.wantErr {
				if err == nil {
					t.Errorf("want error; got none")
				} else {
					//希望出错，结果出错，res有问题直接返回
					return
				}
			}

			if err != nil {
				t.Errorf("error creating trip:%v", err)
				return
			}

			if res.Id != cc.tripID {
				t.Errorf("incorrect id; want%q,got %q", cc.tripID, res.Id)
			}

			b, err := json.Marshal(res.Trip)
			if err != nil {
				t.Errorf("cannot marshall response:%v", err)
			}

			got := string(b)
			if cc.want != got {
				t.Errorf("incorrect response:want %s,got %s", cc.want, got)
			}
		})
	}
}

//profileManager内容测试可以控制
type profileManager struct {
	iID id.IdentityID
	err error
}

//验证身份信息
func (p *profileManager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return p.iID, p.err
}

type carManager struct {
	verifyErr error
	unlockErr error
}

//检查车辆是否被租用
func (c *carManager) Verify(context.Context, id.CarID, *rentalpb.Loaction) error {
	return c.verifyErr
}

//开锁
func (c *carManager) Unlock(context.Context, id.CarID) error {
	return c.unlockErr
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
