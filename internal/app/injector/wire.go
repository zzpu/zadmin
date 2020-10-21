// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"zback/internal/app/api"
	// "zback/internal/app/api/mock"
	"github.com/google/wire"
	"zback/internal/app/bll/impl/bll"
	"zback/internal/app/module/adapter"
	"zback/internal/app/router"

	// mongoModel "zback/internal/app/model/impl/mongo/model"
	gormModel "zback/internal/app/model/impl/gorm/model"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	// 默认使用gorm存储注入，这里可使用 InitMongoDB & mongoModel.ModelSet 替换为 gorm 存储
	wire.Build(
		InitGormDB,
		gormModel.ModelSet, //数据访问层
		// InitMongoDB,
		// mongoModel.ModelSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		bll.BllSet, //业务逻辑层
		api.APISet, //API层
		// mock.MockSet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
