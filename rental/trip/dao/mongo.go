package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	mgutil "coolcar/shared/mongo"

	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	//变量名小写，外面看不到
	col *mongo.Collection
}

//这个文件只知道操纵的是数据库的account表
//操作哪个数据库是auth服务的main函数去确认

//col是操作的表的名称
//newObjID是一个函数，由外部给到，生出一个固定的id
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

//定义表结构
//inline保证数据库里数据层级在外层
type TripRecord struct {
	mgutil.IDField       `bson:"inline"`
	mgutil.UpdateAtField `bson:"inline"`
	Trip                 *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = mgutil.NewObjID()
	r.UpdateAt = mgutil.UpdateAt()

	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
