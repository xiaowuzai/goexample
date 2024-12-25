package dao

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/xiaowuzai/goexample/casbin/service"
	"github.com/xiaowuzai/goexample/db"
)

const (
	defaultTableName = "casbin_rule"
)

type casbinRepo struct {
	e *casbin.Enforcer
}

// NewCasbinRepo  创建casbinRepo.
// confPath： model.conf的路径
// prefix： 数据库前缀可以为空
// tableName： 表名,为空时使用默认的 "casbin_rule"
func NewCasbinRepoByDB(db *db.DB, confPath string, prefix, tableName string) (service.CasbinRepo, error) {
	tn := defaultTableName
	if tableName != "" {
		tn = tableName
	}
	a, err := gormadapter.NewAdapterByDBUseTableName(db.DB, prefix, tn)
	if err != nil {
		return nil, err
	}
	e, err := casbin.NewEnforcer(confPath, a)
	if err != nil {
		return nil, err
	}

	e.LoadPolicy()

	return &casbinRepo{e}, nil
}

// ListPolicy 获取所有的策略
func (c *casbinRepo) ListPolicy() ([][]string, error) {
	return c.e.GetPolicy()
}

// 添加策略
func (c *casbinRepo) AddPolicy(role string, path string, method string) (bool, error) {
	ok, err := c.e.AddPolicy(role, path, method)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	return ok, c.e.LoadPolicy()
}

func (c *casbinRepo) CheckPermission(role string, path string, method string) (bool, error) {
	return c.e.Enforce(role, path, method)
}

func (c *casbinRepo) RemovePolicy(role string, path string, method string) (bool, error) {
	ok, err := c.e.RemovePolicy(role, path, method)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	return ok, c.e.LoadPolicy()
}
