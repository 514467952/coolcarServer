package car

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/shared/id"
)

type Manager struct {
}

//检查车辆是否被租用
func (c *Manager) Verify(context.Context, id.CarID, *rentalpb.Loaction) error {
	return nil
}

//开锁
func (c *Manager) Unlock(context.Context, id.CarID) error {
	return nil
}
