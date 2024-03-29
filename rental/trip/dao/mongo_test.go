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
		mgutil.NewObjectIDWithValue(id.TripID(cc.tripID))

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

func TestGetTrips(t *testing.T) {
	rows := []struct {
		id        string
		accountID string
		status    rentalpb.TripStatus
	}{
		{
			id:        "5f8132eb10714bf629489051",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "5f8132eb10714bf629489052",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "5f8132eb10714bf629489053",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_FINISHED,
		},
		{
			id:        "5f8132eb10714bf629489054",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			id:        "5f8132eb10714bf629489055",
			accountID: "account_id_for_get_trips_1",
			status:    rentalpb.TripStatus_IN_PROGRESS,
		},
	}

	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))

	for _, r := range rows {
		mgutil.NewObjectIDWithValue(id.TripID(r.id))
		_, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountID: r.accountID,
			Status:    r.status,
		})
		if err != nil {
			t.Fatalf("cannot create rows: %v", err)
		}
	}

	cases := []struct {
		name       string
		accountID  string
		status     rentalpb.TripStatus
		wantCount  int
		wantOnlyID string
	}{
		{
			name:      "get_all",
			accountID: "account_id_for_get_trips",
			status:    rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCount: 4,
		},
		{
			name:       "get_in_progress",
			accountID:  "account_id_for_get_trips",
			status:     rentalpb.TripStatus_IN_PROGRESS,
			wantCount:  1,
			wantOnlyID: "5f8132eb10714bf629489054",
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			res, err := m.GetTrips(context.Background(),
				id.AccountID(cc.accountID),
				cc.status)
			if err != nil {
				t.Errorf("cannot get trips: %v", err)
			}

			if cc.wantCount != len(res) {
				t.Errorf("incorrect result count; want: %d, got: %d",
					cc.wantCount, len(res))
			}

			if cc.wantOnlyID != "" && len(res) > 0 {
				if cc.wantOnlyID != res[0].ID.Hex() {
					t.Errorf("only_id incorrect; want: %q, got %q",
						cc.wantOnlyID, res[0].ID.Hex())
				}
			}
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))

	//关心一条记录两个人同时更改的情况
	//固定tripid
	tid := id.TripID("5f8132eb12714bf629489054")
	//构造用户id
	aid := id.AccountID("account_for_update")
	//构造时间
	var now int64 = 10000
	mgutil.NewObjectIDWithValue(tid)
	mgutil.UpdateAt = func() int64 {
		return now
	}
	//创建一个行程
	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountID: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi",
		},
	})

	if err != nil {
		t.Fatalf("cannot create trip:%v", err)
	}

	if tr.UpdateAt != 10000 {
		t.Fatalf("wrong updateat; want: 10000, got:%d", tr.UpdateAt)
	}

	//构造一个更新后的行程
	update := &rentalpb.Trip{
		AccountID: aid.String(),
		Status:    rentalpb.TripStatus_IN_PROGRESS,
		Start: &rentalpb.LocationStatus{
			PoiName: "start_poi_updated",
		},
	}

	cases := []struct {
		name         string
		now          int64
		withUpdateAt int64
		wantErr      bool
	}{
		{
			name:         "normal_update", //正常更新
			now:          20000,           //在20000时更新完成
			withUpdateAt: 10000,           //用20000去更新10000的时间戳
		},
		{
			name:         "update_with_stale_timestamp", //较慢的更新
			now:          30000,
			withUpdateAt: 10000,
			wantErr:      true,
		},
		{
			name:         "update_with_refetch", //重新刷新后读数据
			now:          40000,
			withUpdateAt: 20000, //刷新后的数据应该是normal_update更新后的数据
		},
	}

	for _, cc := range cases {
		now = cc.now
		err := m.UpdateTrip(c, tid, aid, cc.withUpdateAt, update)
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s want error;got none", cc.name)
			} else {
				//测试继续
				continue
			}
		} else {
			if err != nil {
				t.Errorf("%s: cannot update: %v", cc.name, err)
			}
		}

		updatedTrip, err := m.GetTrip(c, tid, aid)
		if err != nil {
			t.Errorf("%s:cannot get trip after update:%v", cc.name, err)
		}

		if cc.now != updatedTrip.UpdateAt {
			t.Errorf("%s:incorrect update: want %d, got:%v", cc.name, cc.now, updatedTrip.UpdateAt)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
