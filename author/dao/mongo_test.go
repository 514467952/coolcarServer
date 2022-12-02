package dao

import (
	"context"
	sharemgo "coolcar/shared/mongo"
	mongotesting "coolcar/shared/testing"
	"fmt"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
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

	//固定id测试，插入几条测试case
	_, err = m.col.InsertMany(c, []interface{}{
		bson.M{
			sharemgo.IDField: mustObjID("63846097e8f95ffe0d631335"),
			openIDField:      "openid_1",
		},
		bson.M{
			sharemgo.IDField: mustObjID("63846097e8f95ffe0d631337"),
			openIDField:      "openid_2",
		},
		bson.M{
			sharemgo.IDField: mustObjID("63846097e8f95ffe0d631336"),
			openIDField:      "openid_3",
		},
	})

	if err != nil {
		t.Fatalf("cannot insert many inital values: %v", err)
	}

	// m.newObjID = func() primitive.ObjectID {
	// 	return mustObjID("63846097e8f95ffe0d63133502")
	// }

	//设置测试case对应的want
	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			want:   "63846097e8f95ffe0d631335",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "63846097e8f95ffe0d631337",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "63846097e8f95ffe0d631336",
		},
	}

	//将测试case每一条的输入输出都对应上
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			id, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("faild resolve account id for %q:%v", cc.openID, err)
			}
			if id != cc.want {
				t.Errorf("resolve account id: want:%q,got:%q", cc.want, id)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}

func mustObjID(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		fmt.Printf("mustObjID error:%v", err)
	}
	return objID
}
