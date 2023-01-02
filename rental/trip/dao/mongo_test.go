package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	mongotesting "coolcar/shared/testing"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//docker启动的mongoDB的端口
var mongoURI string

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	//链接coolcar数据库
	m := NewMongo(mc.Database("coolcar"))

	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountID: "account1",
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
	t.Errorf("inserted row %s with updatedat %v", tr.ID, tr.UpdateAt)
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
