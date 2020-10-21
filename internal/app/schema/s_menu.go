package schema

import (
	"strings"
	"time"

	"zback/pkg/util"
)

// Menu 菜单对象
type Menu struct {
	ID         string    `json:"id"`                                         // 唯一标识
	Title      string    `json:"title" binding:"required"`                   // 菜单名称(显示)
	Component  string    `json:"component" `                                 // 组件
	Name       string    `json:"name" binding:"required"`                    // 菜单名称
	Sequence   int       `json:"sequence"`                                   // 排序值
	Icon       string    `json:"icon"`                                       // 菜单图标
	Path       string    `json:"path"`                                       // 访问路由
	ParentID   string    `json:"parent_id"`                                  // 父级ID
	ParentPath string    `json:"parent_path"`                                // 父级路径
	Level      int       `json:"level"`                                      // 路由层级
	ShowStatus int       `json:"show_status" binding:"required,max=2,min=1"` // 显示状态(1:显示 2:隐藏)
	Status     int       `json:"status" binding:"required,max=2,min=1"`      // 状态(1:启用 2:禁用)
	Memo       string    `json:"memo"`                                       // 备注
	Creator    string    `json:"creator"`                                    // 创建者
	CreatedAt  time.Time `json:"created_at"`                                 // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`                                 // 更新时间
	Method     string    `yaml:"method" json:"method"`                       // 按钮类型(其实是HTTP的方法,GET,POST,PUT,DELETE等)
}

func (a *Menu) String() string {
	return util.JSONMarshalToString(a)
}

// MenuQueryParam 查询条件
type MenuQueryParam struct {
	PaginationParam
	IDs              []string `form:"-"`          // 唯一标识列表
	Name             string   `form:"-"`          // 菜单名称
	PrefixParentPath string   `form:"-"`          // 父级路径(前缀模糊查询)
	QueryValue       string   `form:"queryValue"` // 模糊查询
	ParentID         *string  `form:"parentID"`   // 父级内码
	Level            int      `form:"level"`      // 菜单层级(3级为按钮)
	ShowStatus       int      `form:"showStatus"` // 显示状态(1:显示 2:隐藏)
	Status           int      `form:"status"`     // 状态(1:启用 2:禁用)
}

// MenuQueryOptions 查询可选参数项
type MenuQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuQueryResult 查询结果
type MenuQueryResult struct {
	Data       Menus
	PageResult *PaginationResult
}

// Menus 菜单列表
type Menus []*Menu

func (a Menus) Len() int {
	return len(a)
}

func (a Menus) Less(i, j int) bool {
	return a[i].Sequence > a[j].Sequence
}

func (a Menus) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ToMap 转换为键值映射
func (a Menus) ToMap() map[string]*Menu {
	m := make(map[string]*Menu)
	for _, item := range a {
		m[item.ID] = item
	}
	return m
}

func (a Menus) ToParentIDMap() map[string]*Menu {
	m := make(map[string]*Menu)
	for _, item := range a {
		m[item.ParentID] = item
	}
	return m
}

// SplitParentIDs 拆分父级路径的唯一标识列表
func (a Menus) SplitParentIDs() []string {
	idList := make([]string, 0, len(a))
	mIDList := make(map[string]struct{})

	for _, item := range a {
		if _, ok := mIDList[item.ID]; ok || item.ParentPath == "" {
			continue
		}

		for _, pp := range strings.Split(item.ParentPath, "/") {
			if _, ok := mIDList[pp]; ok {
				continue
			}
			idList = append(idList, pp)
			mIDList[pp] = struct{}{}
		}
	}

	return idList
}

// ToTree 转换为菜单树
func (a Menus) ToTree() []*MenuTree {
	list := make([]*MenuTree, len(a))
	for i, item := range a {
		list[i] = &MenuTree{
			ID:         item.ID,
			Title:      item.Title,
			Name:       item.Name,
			Icon:       item.Icon,
			Path:       item.Path,
			ParentID:   item.ParentID,
			ParentPath: item.ParentPath,
			Sequence:   item.Sequence,
			ShowStatus: item.ShowStatus,
			Status:     item.Status,
			Level:      item.Level,
			Component:  item.Component,
			Method:     item.Method,
		}
	}
	mi := make(map[string]*MenuTree)
	for _, tr := range list {
		tr.Children = make([]*MenuTree, 0)
		mi[tr.ID] = tr
	}

	var trs []*MenuTree
	for _, tr := range list {
		if tr.ParentID == "" {
			trs = append(trs, tr)
			continue
		}
		if pitem, ok := mi[tr.ParentID]; ok {
			pitem.Children = append(pitem.Children, tr)
		}
	}
	return trs

}

// ----------------------------------------MenuTree--------------------------------------

// MenuTree 菜单树
type MenuTree struct {
	ID         string      `yaml:"-" json:"id"`                        // 唯一标识
	Title      string      `yaml:"title" json:"title"`                 // 菜单名称（显示)
	Name       string      `yaml:"name" json:"name"`                   // 菜单名称
	Component  string      `yaml:"component" json:"component"`         // 组件
	Icon       string      `yaml:"icon" json:"icon"`                   // 菜单图标
	Path       string      `yaml:"path,omitempty" json:"path"`         // 访问路由
	ParentID   string      `yaml:"-" json:"parent_id"`                 // 父级ID
	ParentPath string      `yaml:"-" json:"parent_path"`               // 父级路径
	Sequence   int         `yaml:"sequence" json:"sequence"`           // 排序值
	Level      int         `yaml:"level" json:"level"`                 // 路由层级
	ShowStatus int         `yaml:"-" json:"show_status"`               // 显示状态(1:显示 2:隐藏)
	Status     int         `yaml:"-" json:"status"`                    // 状态(1:启用 2:禁用)
	Method     string      `yaml:"method" json:"method"`               // 按钮类型
	Children   []*MenuTree `yaml:"children,omitempty" json:"children"` // 子级树
}

//// ToTree 转换为树形结构
//func (a []*MenuTree) ToTree() []*MenuTree {
//
//}

// ----------------------------------------MenuAction--------------------------------------

// MenuAction 菜单动作对象
type MenuAction struct {
	ID        string        `yaml:"-" json:"id"`                           // 唯一标识
	MenuID    string        `yaml:"-" binding:"required" json:"menu_id"`   // 菜单ID
	Code      string        `yaml:"code" binding:"required" json:"code"`   // 动作编号
	Title     string        `yaml:"title" binding:"required" json:"title"` // 动作名称
	Name      string        `yaml:"name" binding:"required" json:"name"`   // 动作名称
	Resources MenuResources `yaml:"resources,omitempty" json:"resources"`  // 资源列表
}

// MenuActionQueryParam 查询条件
type MenuActionQueryParam struct {
	PaginationParam
	MenuID string   // 菜单ID
	IDs    []string // 唯一标识列表
}

// MenuActionQueryOptions 查询可选参数项
type MenuActionQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuActionQueryResult 查询结果
type MenuActionQueryResult struct {
	Data       MenuActions
	PageResult *PaginationResult
}

// MenuActions 菜单动作管理列表
type MenuActions []*MenuAction

// ToMap 转换为map
func (a MenuActions) ToMap() map[string]*MenuAction {
	m := make(map[string]*MenuAction)
	for _, item := range a {
		m[item.Code] = item
	}
	return m
}

// FillResources 填充资源数据
func (a MenuActions) FillResources(mResources map[string]MenuResources) {
	for i, item := range a {
		a[i].Resources = mResources[item.ID]
	}
}

// ToMenuIDMap 转换为菜单ID映射
func (a MenuActions) ToMenuIDMap() map[string]MenuActions {
	m := make(map[string]MenuActions)
	for _, item := range a {
		m[item.MenuID] = append(m[item.MenuID], item)
	}
	return m
}

// ----------------------------------------MenuResource--------------------------------------

// MenuResource 菜单动作关联资源对象
type MenuResource struct {
	ID     string `yaml:"-" json:"id"`                             // 唯一标识
	MenuID string `yaml:"-" json:"action_id"`                      // 菜单动作ID
	Method string `yaml:"method" binding:"required" json:"method"` // 资源请求方式(支持正则)
	Path   string `yaml:"path" binding:"required" json:"path"`     // 资源请求路径（支持/:id匹配）
}

// MenuResourceQueryParam 查询条件
type MenuResourceQueryParam struct {
	PaginationParam
	MenuID  string   // 菜单ID
	MenuIDs []string // 菜单ID列表
}

// MenuResourceQueryOptions 查询可选参数项
type MenuResourceQueryOptions struct {
	OrderFields []*OrderField // 排序字段
}

// MenuResourceQueryResult 查询结果
type MenuResourceQueryResult struct {
	Data       MenuResources
	PageResult *PaginationResult
}

// MenuResources 菜单动作关联资源管理列表
type MenuResources []*MenuResource

// ToMap 转换为map
func (a MenuResources) ToMap() map[string]*MenuResource {
	m := make(map[string]*MenuResource)
	for _, item := range a {
		m[item.Method+item.Path] = item
	}
	return m
}

// ToMenuIDMap 转换为动作ID映射
func (a MenuResources) ToMenuIDMap() map[string]MenuResources {
	m := make(map[string]MenuResources)
	for _, item := range a {
		m[item.MenuID] = append(m[item.MenuID], item)
	}
	return m
}
