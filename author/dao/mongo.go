package dao

import (
	"context"
	sharemgo "coolcar/shared/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"
const IDField = "_id"

type MyMongo struct {
	//变量名小写，外面看不到
	col *mongo.Collection
}

//这个文件只知道操纵的是数据库的account表
//操作哪个数据库是auth服务的main函数去确认
func NewMongo(db *mongo.Database) *MyMongo {
	return &MyMongo{
		col: db.Collection("account"),
	}
}

//收到openID，返回对应的记录ID
func (m *MyMongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: openID,
	}, sharemgo.Set(bson.M{
		openIDField: openID,
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
