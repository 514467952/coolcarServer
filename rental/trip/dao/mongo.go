package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField      = "trip"
	accountIDField = tripField + ".accountid"
	statusField    = tripField + ".status"
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

//同一个account最多只能有一个进行中的Trip
//强类型化Tripid
//表格驱动测试

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

//行程服务的获取
func (m *Mongo) GetTrip(c context.Context, id id.TripID, accountID id.AccountID) (*TripRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("GetTrip id无效:%v", err)
	}
	res := m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
		accountIDField:     accountID,
	})
	if err := res.Err(); err != nil {
		return nil, err
	}
	var tr TripRecord
	err = res.Decode(&tr)
	if err != nil {
		return nil, fmt.Errorf("GetTrip cannot decode:%v", err)
	}
	return &tr, nil
}

//批量获取行程
func (m *Mongo) GetTrips(c context.Context, accountID id.AccountID, status rentalpb.TripStatus) ([]*TripRecord, error) {
	//通过accountID用户的全部行程
	filter := bson.M{
		accountIDField: accountID.String(),
	}

	//如果外部给了status，再加上行程的状态
	if status != rentalpb.TripStatus_TS_NOT_SPECIFIED {
		filter[statusField] = status
	}

	//调用find，第一个参数是context，第二个是查询条件
	//res类型是是mongo.Cursor
	res, err := m.col.Find(c, filter)
	if err != nil {
		return nil, err
	}

	var trips []*TripRecord
	for res.Next(c) {
		var trip TripRecord
		err := res.Decode(&trip)
		//如果其中一行出错，暂时定义为全部出错
		if err != nil {
			return nil, err
		}

		trips = append(trips, &trip)
	}
	return trips, nil
}

//更新行程
//用updatedAt实现一个乐观锁来解决同时更新的问题
func (m *Mongo) UpdateTrip(c context.Context, tid id.TripID, aid id.AccountID, updatedAt int64, trip *rentalpb.Trip) error {
	objID, err := objid.FromID(tid)
	if err != nil {
		return fmt.Errorf("invalid id:%v", err)
	}

	newUpdatedAt := mgutil.UpdateAt()
	res, err := m.col.UpdateOne(c, bson.M{
		mgutil.IDFieldName:        objID,
		accountIDField:            aid.String(),
		mgutil.UpdatedAtFieldName: updatedAt,
	}, mgutil.Set(bson.M{
		tripField:                 trip,
		mgutil.UpdatedAtFieldName: newUpdatedAt,
	}))

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		//没有匹配任何的文档
		return mongo.ErrNoDocuments
	}
	return nil
}
