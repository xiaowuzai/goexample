package server

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaowuzai/goexample/casbin/middleware"
	"github.com/xiaowuzai/goexample/casbin/service"
)

type Server struct {
	service *service.CasbinService
	router  *gin.Engine
}

func NewServer(s *service.CasbinService) *Server {
	server := &Server{
		service: s,
	}

	server.setRouter()
	return server
}

func (s *Server) setRouter() {
	e := gin.Default()

	// 初始化 Casbin 中间件
	casbinMiddleware := middleware.CasbinMiddleware(s.service)

	v1 := e.Group("/v1")
	{
		v1.POST("/policies", s.AddPolicy)

		v1.Use(middleware.AuthMiddleware())

		// 应用 Casbin 中间件到需要权限控制的路由
		v1.Use(casbinMiddleware)

		// 定义策略相关的路由
		v1.GET("/policies", s.GetPolicies)
		v1.DELETE("/policies", s.DeletePolicy)

		// 测试路由
		v1.GET("/test", s.Test)
		v1.GET("/bill", s.Bill)
	}

	s.router = e
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
