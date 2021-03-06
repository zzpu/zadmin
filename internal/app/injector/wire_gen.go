// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"zback/internal/app/api"
	"zback/internal/app/bll/impl/bll"
	"zback/internal/app/model/impl/gorm/model"
	"zback/internal/app/module/adapter"
	"zback/internal/app/router"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	auther, cleanup, err := InitAuth()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := InitGormDB()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	role := &model.Role{
		DB: db,
	}
	roleMenu := &model.RoleMenu{
		DB: db,
	}
	menu := &model.Menu{
		DB: db,
	}
	user := &model.User{
		DB: db,
	}
	userRole := &model.UserRole{
		DB: db,
	}
	casbinAdapter := &adapter.CasbinAdapter{
		RoleModel:         role,
		RoleMenuModel:     roleMenu,
		MenuResourceModel: menu,
		UserModel:         user,
		UserRoleModel:     userRole,
	}
	syncedEnforcer, cleanup3, err := InitCasbin(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	demo := &model.Demo{
		DB: db,
	}
	bllDemo := &bll.Demo{
		DemoModel: demo,
	}
	apiDemo := &api.Demo{
		DemoBll: bllDemo,
	}
	login := &bll.Login{
		Auth:          auther,
		UserModel:     user,
		UserRoleModel: userRole,
		RoleModel:     role,
		RoleMenuModel: roleMenu,
		MenuModel:     menu,
	}
	apiLogin := &api.Login{
		LoginBll: login,
	}
	trans := &model.Trans{
		DB: db,
	}
	bllMenu := &bll.Menu{
		TransModel: trans,
		MenuModel:  menu,
	}
	apiMenu := &api.Menu{
		MenuBll: bllMenu,
	}
	bllRole := &bll.Role{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		RoleModel:     role,
		RoleMenuModel: roleMenu,
		UserModel:     user,
	}
	apiRole := &api.Role{
		RoleBll: bllRole,
	}
	bllUser := &bll.User{
		Enforcer:      syncedEnforcer,
		TransModel:    trans,
		UserModel:     user,
		UserRoleModel: userRole,
		RoleModel:     role,
	}
	apiUser := &api.User{
		UserBll: bllUser,
	}
	routerRouter := &router.Router{
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		DemoAPI:        apiDemo,
		LoginAPI:       apiLogin,
		MenuAPI:        apiMenu,
		RoleAPI:        apiRole,
		UserAPI:        apiUser,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine:         engine,
		Auth:           auther,
		CasbinEnforcer: syncedEnforcer,
		MenuBll:        bllMenu,
	}
	return injector, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
