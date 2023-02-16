package profile

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/shared/id"
	"encoding/base64"
	"fmt"

	"github.com/gogo/protobuf/proto"
)

type Fetcher interface {
	GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error)
}

type Manager struct {
	Fetcher Fetcher
}

//认证方法:通过Verify函数告诉我们是否有资格
//返回IdentityID 租车时，可能会重新审查，换身份信息
func (m *Manager) Verify(c context.Context, aid id.AccountID) (id.IdentityID, error) {
	nilID := id.IdentityID("")

	p, err := m.Fetcher.GetProfile(c, &rentalpb.GetProfileRequest{})
	if err != nil {
		return nilID, fmt.Errorf("Verify cannot get profile:%v", err)
	}

	if p.IdentityStatus != rentalpb.IdentityStatus_VERIFIED {
		return nilID, fmt.Errorf("Verify invalid identity status")
	}

	b, err := proto.Marshal(p.Identity)
	if err != nil {
		return nilID, fmt.Errorf("Verify cannot marshal identity:%v", err)
	}

	return id.IdentityID(base64.StdEncoding.EncodeToString(b)), nil

}
