package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaowuzai/goexample/casbin/service"
)

func CasbinMiddleware(casbin service.CasbinRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求上下文中获取用户名和角色，例如通过JWT解析
		username, exists := c.Get("username")
		if !exists {
			log.Printf("未找到用户名")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未授权"})
			c.Abort()
			return
		}

		roleNames, exists := c.Get("roles")
		if !exists {
			log.Printf("未找到用户角色")
			c.JSON(http.StatusForbidden, gin.H{"message": "权限不足"})
			c.Abort()
			return
		}

		isAdmin, exists := c.Get("isAdmin")
		if !exists {
			isAdmin = false
		}

		if isAdmin.(bool) {
			c.Next()
			return
		}
		obj := c.Request.URL.Path // 目标对象，通常是请求的 URL
		act := c.Request.Method   // 动作，通常是 HTTP 方法
		for _, role := range roleNames.([]string) {
			// 检查权限
			ok, err := casbin.CheckPermission(role, obj, act)
			if err != nil {
				log.Printf("检查权限失败: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "检查权限失败"})
				c.Abort()
				return
			}
			if ok {
				log.Printf("用户(%s) %s方式访问资源 %s, 权限通过", username, c.Request.Method, c.Request.URL.String())
				// 如果通过了权限检查，则继续处理请求
				c.Next()
				return
			}
		}

		log.Printf("用户(%s) %s方式访问资源 %s, 权限不足.",
			username, c.Request.Method, c.Request.URL.String())

		c.JSON(http.StatusForbidden, gin.H{"message": "权限不足"})
		c.Abort()
	}
}
