package profile

import (
	"context"
	"coolcar/shared/id"
)

type Manager struct {
}

//认证方法:通过Verify函数告诉我们是否有资格
//返回IdentityID 租车时，可能会重新审查，换身份信息
func (p *Manager) Verify(context.Context, id.AccountID) (id.IdentityID, error) {
	return id.IdentityID("identity1"), nil
}
