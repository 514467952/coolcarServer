/*练习使用go语言操作mongoDB*/
package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func main() {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://106.54.49.241:27017/?readPreference=primary&ssl=false&directConnection=true"))
	if err != nil {
		panic(err)
	}
	col := mc.Database("coolcar").Collection("account")
	// findRows(c, col)
	findAll(c, col)
}

//查找全部
func findAll(c context.Context, col *mongo.Collection) {
	cur, err := col.Find(c, bson.M{})
	if err != nil {
		panic(err)
	}
	for cur.Next(c) {
		var row struct {
			ID     primitive.ObjectID `bson:"_id"`
			OpenID string             `bson:"open_id"`
		}
		err = cur.Decode(&row)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", row)
	}
}

//查找单个
func findRows(c context.Context, col *mongo.Collection) {
	res := col.FindOne(c, bson.M{
		"open_id": "123",
	})
	// fmt.Printf("%+v\n", res)
	var row struct {
		ID     primitive.ObjectID `bson:"_id"`
		OpenID string             `bson:"open_id"`
	}
	err := res.Decode(&row)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", row)
}

//插入记录
func insertRow(c context.Context, col *mongo.Collection) {
	res, err := col.InsertMany(c, []interface{}{
		bson.M{
			"open_id": "123",
		},
		bson.M{
			"open_id": "456",
		},
	})
	if err != nil {
		panic(err)
	}

	//%+v会把res字段名+值格式化打印
	fmt.Printf("%+v", res)
}
