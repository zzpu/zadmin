package model

import "github.com/google/wire"

// ModelSet model注入
var ModelSet = wire.NewSet(
	DemoSet,
	MenuResourceSet,
	MenuActionSet,
	MenuSet,
	RoleMenuSet,
	RoleSet,
	TransSet,
	UserRoleSet,
	UserSet,
)
