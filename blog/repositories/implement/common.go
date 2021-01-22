package implement

import (
	"blog/repositories/common"
	"blog/repositories/interfaces"
)

var (
	baseRepo interfaces.BaseRepositoryInterface = common.NewMysqlRepository()
)
