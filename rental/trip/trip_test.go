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
	"math/rand"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()

	pm := &profileManager{}
	cm := &carManager{}
	s := newService(c, t, pm, cm)

	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Loaction{
			Latitude:  32.123,
			Longitude: 114.2525,
		},
	}
	pm.iID = "identity1"
	golden := `{"accountID":%q,"carID":"car1","start":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":1675645239},"current":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":1675645239},"status":1,"identity_id":"identity1"}`
	nowFunc = func() int64 {
		return 1675645239
	}
	cases := []struct {
		name         string
		accountID    string
		tripID       string
		profileErr   error
		carVerifyErr error
		carUnlockErr error
		want         string
		wantErr      bool
	}{
		{
			name:      "normal_create",
			accountID: "account1",
			tripID:    "5f8132eb00714bf62948905c",
			want:      fmt.Sprintf(golden, "account1"),
		},
		{
			name:       "profile_err",
			accountID:  "account2",
			tripID:     "5f8132eb00714bf629489055",
			profileErr: fmt.Errorf("profile"),
			wantErr:    true,
		},
		{
			name:         "car_verify_err",
			accountID:    "account3",
			tripID:       "5f8132eb00714bf629489056",
			carVerifyErr: fmt.Errorf("verify"),
			wantErr:      true,
		},
		{
			name:         "car_unlock_err", //开锁失败，但是行程创建成功
			accountID:    "account4",
			tripID:       "5f8132eb00714bf629489057",
			carUnlockErr: fmt.Errorf("unlock"),
			want:         fmt.Sprintf(golden, "account4"),
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			//tripid
			mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))

			pm.err = cc.profileErr
			cm.unlockErr = cc.carUnlockErr
			cm.verifyErr = cc.carVerifyErr
			c := sharedauth.ContextWithAccountID(
				context.Background(), id.AccountID(cc.accountID))

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

//测试行程服务完整生命周期
func TestTripLifecycle(t *testing.T) {
	c := sharedauth.ContextWithAccountID(
		context.Background(),
		id.AccountID("account_for_lifecycle"),
	)

	s := newService(c, t, &profileManager{}, &carManager{})

	tid := id.TripID("5f8132eb07714bf629489057")
	mgutil.NewObjectIDWithValue(tid)
	cases := []struct {
		name string
		now  int64
		op   func() (*rentalpb.Trip, error)
		want string
	}{
		{
			name: "create_trip",
			now:  10000,
			op: func() (*rentalpb.Trip, error) {
				e, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
					CarId: "car1",
					Start: &rentalpb.Loaction{
						Latitude:  32.123,
						Longitude: 114.2525,
					},
				})
				if err != nil {
					return nil, err
				}
				return e.Trip, nil
			},
			want: `{"accountID":"account_for_lifecycle","carID":"car1","start":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"status":1}`,
		},
		{
			name: "update_trip",
			now:  20000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
					Current: &rentalpb.Loaction{
						Latitude:  28.234234,
						Longitude: 123.243255,
					},
				})
			},
			want: `{"accountID":"account_for_lifecycle","carID":"car1","start":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"Loaction":{"latitude":28.234234,"longitude":123.243255},"fee_cent":3857,"km_driven":305.0839256834663,"poi_name":"中关村","timestamp_sec":20000},"status":1}`,
		},
		{
			name: "finish_trip",
			now:  30000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id:      tid.String(),
					EndTrip: true,
				})
			},
			want: `{"accountID":"account_for_lifecycle","carID":"car1","start":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"Loaction":{"latitude":28.234234,"longitude":123.243255},"fee_cent":7542,"km_driven":538.6937581015429,"poi_name":"中关村","timestamp_sec":30000},"end":{"Loaction":{"latitude":28.234234,"longitude":123.243255},"fee_cent":7542,"km_driven":538.6937581015429,"poi_name":"中关村","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "query_trip",
			now:  40000,
			op: func() (*rentalpb.Trip, error) {
				return s.GetTrip(c, &rentalpb.GetTripRequest{
					Id: tid.String(),
				})
			},
			want: `{"accountID":"account_for_lifecycle","carID":"car1","start":{"Loaction":{"latitude":32.123,"longitude":114.2525},"poi_name":"天安门","timestamp_sec":10000},"current":{"Loaction":{"latitude":28.234234,"longitude":123.243255},"fee_cent":7542,"km_driven":538.6937581015429,"poi_name":"中关村","timestamp_sec":30000},"end":{"Loaction":{"latitude":28.234234,"longitude":123.243255},"fee_cent":7542,"km_driven":538.6937581015429,"poi_name":"中关村","timestamp_sec":30000},"status":2}`,
		},
	}

	rand.Seed(1345)
	for _, cc := range cases {
		//固定时间
		nowFunc = func() int64 {
			return cc.now
		}
		trip, err := cc.op()
		if err != nil {
			t.Errorf("%s:operation failed:%v", cc.name, err)
			continue
		}
		b, err := json.Marshal(trip)
		if err != nil {
			t.Errorf("%s: failed marshalling response:%v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			t.Errorf("%s:incorrect response;want:%s,got:%s", cc.name, cc.want, got)
		}
	}
}

//创建数据库链接方法单独提出来
func newService(c context.Context, t *testing.T, pm ProfileManager, cm CarManager) *Service {
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create mongo client: %v", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("cannot create logger:%v", err)
	}

	db := mc.Database("coolcar")
	mongotesting.SetupIndexes(c, db)
	return &Service{
		ProfileManager: pm,
		CarManager:     cm,
		POIManager:     &poi.Manager{},
		Mongo:          dao.NewMongo(db),
		Logger:         logger,
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
