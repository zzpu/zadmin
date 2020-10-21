package entity

import (
	"context"

	"github.com/jinzhu/gorm"
	"zback/internal/app/schema"
	"zback/pkg/util"
)

// GetMenuDB 获取菜单存储
func GetMenuDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return GetDBWithModel(ctx, defDB, new(Menu))
}

// SchemaMenu 菜单对象
type SchemaMenu schema.Menu

// ToMenu 转换为菜单实体
func (a SchemaMenu) ToMenu() *Menu {
	item := new(Menu)
	util.StructMapToStruct(a, item)
	return item
}

// Menu 菜单实体
type Menu struct {
	Model
	Title      string  `gorm:"column:title;size:50;index;default:'';not null;"`     // 菜单名称(显示)
	Name       string  `gorm:"column:name;size:50;index;default:'';not null;"`      // 菜单名称
	Component  string  `gorm:"column:component;size:50;index;default:'';not null;"` // 组件
	Method     string  `gorm:"column:method;size:50;index;default:'';"`             // 方法,GET,POST等,只有按钮资源有
	Sequence   int     `gorm:"column:sequence;index;default:0;not null;"`           // 排序值
	Icon       *string `gorm:"column:icon;size:255;"`                               // 菜单图标
	Path       *string `gorm:"column:path;size:255;"`                               // 访问路由
	ParentID   *string `gorm:"column:parent_id;size:36;index;"`                     // 父级内码
	ParentPath *string `gorm:"column:parent_path;size:518;index;"`                  // 父级路径
	Level      int     `gorm:"column:level;index;default:0;not null;"`              // 路由层级
	ShowStatus int     `gorm:"column:show_status;index;default:0;not null;"`        // 状态(1:显示 2:隐藏)
	Status     int     `gorm:"column:status;index;default:0;not null;"`             // 状态(1:启用 2:禁用)
	Memo       *string `gorm:"column:memo;size:1024;"`                              // 备注
	Creator    string  `gorm:"column:creator;size:36;"`                             // 创建人
}

// TableName 表名
func (a Menu) TableName() string {
	return a.Model.TableName("menu")
}

// ToSchemaMenu 转换为菜单对象
func (a Menu) ToSchemaMenu() *schema.Menu {
	item := new(schema.Menu)
	util.StructMapToStruct(a, item)
	return item
}

// Menus 菜单实体列表
type Menus []*Menu

// ToSchemaMenus 转换为菜单对象列表
func (a Menus) ToSchemaMenus() []*schema.Menu {
	list := make([]*schema.Menu, len(a))
	for i, item := range a {
		//数据库的菜单对象转化为表单的菜单对象
		list[i] = item.ToSchemaMenu()
	}
	return list
}
