package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Bill(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "bill",
	})
}

func (s *Server) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

// GetPolicies 处理获取所有策略的请求
func (s *Server) GetPolicies(c *gin.Context) {
	policies, err := s.service.ListPolicy()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取策略失败"})
		return
	}
	c.JSON(http.StatusOK, policies)
}

// AddPolicy 处理添加策略的请求
func (s *Server) AddPolicy(c *gin.Context) {
	var policy struct {
		Role   string `json:"role" binding:"required"`
		Path   string `json:"path" binding:"required"`
		Method string `json:"method" binding:"required"`
	}

	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求参数"})
		return
	}

	ok, err := s.service.AddPolicy(policy.Role, policy.Path, policy.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "添加策略失败"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "策略已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "策略添加成功"})
}

// DeletePolicy 处理删除策略的请求
func (s *Server) DeletePolicy(c *gin.Context) {
	var policy struct {
		Role   string `json:"role" binding:"required"`
		Path   string `json:"path" binding:"required"`
		Method string `json:"method" binding:"required"`
	}

	if err := c.ShouldBindJSON(&policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无效的请求参数"})
		return
	}

	ok, err := s.service.RemovePolicy(policy.Role, policy.Path, policy.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "删除策略失败"})
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "策略不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "策略删除成功"})
}
