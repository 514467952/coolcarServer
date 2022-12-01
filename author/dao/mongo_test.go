package dao

import (
	"context"
	mongotesting "coolcar/shared/testing"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//docker启动的mongoDB的端口
var mongoURI string

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	//链接coolcar数据库
	m := NewMongo(mc.Database("coolcar"))

	m.newObjID = func() primitive.ObjectID {
		objID, _ := primitive.ObjectIDFromHex("63846097e8f95ffe0d631335")
		return objID
	}

	//调用获取accountID的函数，传入open_id为123
	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Errorf("faild resolve account id for 123:%v", err)
	} else {
		//得到的值不对，要报错
		//%q是在输出的字符串两边加上""
		want := "63846097e8f95ffe0d631335"
		if id != want {
			t.Errorf("resolve account id: want:%q,got:%q", want, id)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
