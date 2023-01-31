package poi

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"hash/fnv"

	"github.com/gogo/protobuf/proto"
)

var poi = []string{
	"中关村",
	"天安门",
	"陆家嘴",
	"迪士尼",
	"环球影城",
	"鸟巢",
}

type Manager struct {
}

func (*Manager) Resolve(c context.Context, loc *rentalpb.Loaction) (string, error) {
	//将req转为二进制数组
	b, err := proto.Marshal(loc)
	if err != nil {
		return "", err
	}
	//将req转为一个哈希值
	h := fnv.New32()
	h.Write(b)

	//给定一个loc，只会算出一个哈希值
	return poi[int(h.Sum32())%len(poi)], nil
}
