package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/shared/id"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/testing"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

//docker启动的mongoDB的端口
// var mongoURI string

func TestGetTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	//链接coolcar数据库
	m := NewMongo(mc.Database("coolcar"))
	acct := id.AccountID("account2")
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
		Status: rentalpb.TripStatus_IN_PROGRESS,
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
