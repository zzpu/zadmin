package bll

import (
	"context"

	"github.com/casbin/casbin/v2"
	"zback/internal/app/config"
	"zback/pkg/logger"
)

var chCasbinPolicy chan *chCasbinPolicyItem

type chCasbinPolicyItem struct {
	ctx context.Context
	e   *casbin.SyncedEnforcer
}

func init() {
	chCasbinPolicy = make(chan *chCasbinPolicyItem, 1)
	go func() {
		//这里不断读取管道,当有信号过来就加载策略 -->保证每次关于策略(用户,角色,菜单)的更新都会及时加载
		for item := range chCasbinPolicy {
			err := item.e.LoadPolicy()
			if err != nil {
				logger.Errorf(item.ctx, "The load casbin policy error: %s", err.Error())
			}
		}
	}()
}

/**
 * 异步加载casbin权限策略
 *
 * param: context.Context        ctx
 * param: *casbin.SyncedEnforcer e
 */
func LoadCasbinPolicy(ctx context.Context, e *casbin.SyncedEnforcer) {
	if !config.C.Casbin.Enable {
		return
	}

	if len(chCasbinPolicy) > 0 {
		logger.Infof(ctx, "The load casbin policy is already in the wait queue")
		return
	}

	chCasbinPolicy <- &chCasbinPolicyItem{
		ctx: ctx,
		e:   e,
	}
}
