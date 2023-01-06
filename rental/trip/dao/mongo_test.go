package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/testing"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
)

//docker启动的mongoDB的端口
// var mongoURI string

//创建行程测试
func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	//链接coolcar数据库
	db := mc.Database("coolcar")
	err = mongotesting.SetupIndexes(c, db)
	if err != nil {
		t.Fatalf("cannot setup indexes: %v", err)
	}

	m := NewMongo(db)

	cases := []struct {
		name       string
		tripID     string
		accountID  string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finished", //行程完成1
			tripID:     "5f8132eb00714bf62948905c",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "another_finished", //行程完成2
			tripID:     "5f8132eb00714bf62948905d",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "in_progress", //进行中的行程
			tripID:     "5f8132eb00714bf62948905e",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			name:       "another_in_progress", //进行中的行程2
			tripID:     "5f8132eb00714bf62948905f",
			accountID:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "in_progress_by_another_account", //另外一个用户进行中的行程
			tripID:     "5f8132eb00714bf629489060",
			accountID:  "account2",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	//t.Run()不能保证执行case时的顺序
	for _, cc := range cases {
		mgutil.NewObjID = func() primitive.ObjectID {
			return objid.MustFromID(id.TripID(cc.tripID))
		}

		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountID: cc.accountID,
			Status:    cc.tripStatus,
		})

		//预期结果出错，但是没有
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s want err but no error: got none", cc.name)
			}
			continue
		}

		if err != nil {
			t.Errorf("%s error creating trip:%v", cc.name, err)
			continue
		}

		if tr.ID.Hex() != cc.tripID {
			t.Errorf("%s incorrect trip id:want:%q;got:%q", cc.name, cc.tripID, tr.ID.Hex())
		}

	}
}

//获取行程测试
func TestGetTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	//链接coolcar数据库
	m := NewMongo(mc.Database("coolcar"))
	acct := id.AccountID("account1")
	//确保先测试创建行程后，获取行程不受影响
	mgutil.NewObjID = primitive.NewObjectID
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountID: acct.String(),
		CarID:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "startpoint",
			Loaction: &rentalpb.Loaction{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName:  "endpoint",
			FeeCent:  10000,
			KmDriven: 35,
			Loaction: &rentalpb.Loaction{
				Latitude:  35,
				Longitude: 115,
			},
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Errorf("cannot create trip:%v", err)
	}

	got, err := m.GetTrip(c, objid.ToTripID(tr.ID), acct)
	if err != nil {
		t.Fatalf("cannot get trip:%v", err)
	}

	if diff := cmp.Diff(tr, got, protocmp.Transform()); diff != "" {
		t.Errorf("result differs:-want +got: %s", diff)
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
