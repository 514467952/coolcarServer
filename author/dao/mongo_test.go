package dao

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://106.54.49.241:27017/?readPreference=primary&ssl=false&directConnection=true"))
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}
	//链接coolcar数据库
	m := NewMongo(mc.Database("coolcar"))
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
