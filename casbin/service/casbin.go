package service

// PolicyRepo casbin policy 的数据库操作
type CasbinRepo interface {
	ListPolicy() ([][]string, error)
	AddPolicy(role string, path string, method string) (bool, error)
	CheckPermission(role string, path string, method string) (bool, error)
	RemovePolicy(role string, path string, method string) (bool, error)
}

type CasbinService struct {
	repo CasbinRepo
}

func NewCasbinService(repo CasbinRepo) *CasbinService {
	return &CasbinService{
		repo: repo,
	}
}

func (c *CasbinService) ListPolicy() ([][]string, error) {
	return c.repo.ListPolicy()
}
func (c *CasbinService) AddPolicy(role string, path string, method string) (bool, error) {
	return c.repo.AddPolicy(role, path, method)
}
func (c *CasbinService) CheckPermission(role string, path string, method string) (bool, error) {
	return c.repo.CheckPermission(role, path, method)
}
func (c *CasbinService) RemovePolicy(role string, path string, method string) (bool, error) {
	return c.repo.RemovePolicy(role, path, method)
}
