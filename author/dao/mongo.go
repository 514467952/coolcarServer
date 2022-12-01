package dao

import (
	"context"
	sharemgo "coolcar/shared/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

type MyMongo struct {
	//变量名小写，外面看不到
	col      *mongo.Collection
	newObjID func() primitive.ObjectID
}

//这个文件只知道操纵的是数据库的account表
//操作哪个数据库是auth服务的main函数去确认

//col是操作的表的名称
//newObjID是一个函数，由外部给到，生出一个固定的id
func NewMongo(db *mongo.Database) *MyMongo {
	return &MyMongo{
		col:      db.Collection("account"),
		newObjID: primitive.NewObjectID,
	}
}

//收到openID，返回对应的记录ID
func (m *MyMongo) ResolveAccountID(c context.Context, openID string) (string, error) {

	//生成一个固定ID
	insertedID := m.newObjID()
	//先去查找记录，如果找到直接返回
	//找不到，插入一条固定ID：openID的记录
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: openID,
	}, sharemgo.SetOnInsert(bson.M{
		sharemgo.IDField: insertedID,
		openIDField:      openID,
	}), options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))

	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate:%v", err)
	}

	var row sharemgo.ObjectID

	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode result:%v", err)
	}

	return row.ID.Hex(), nil
}
