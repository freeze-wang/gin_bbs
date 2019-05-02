package routes

import (
	"gin_bbs/pkg/ginutils/csrf"
	"gin_bbs/pkg/ginutils/oldvalue"
	"gin_bbs/pkg/ginutils/session"

	"gin_bbs/pkg/ginutils/router"

	"github.com/gin-gonic/gin"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {
	// ---------------------------------- 注册全局中间件 ----------------------------------
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	// 自定义全局中间件
	g.Use(session.SessionMiddleware())   // session
	g.Use(csrf.CsrfMiddleware())         // csrf
	g.Use(oldvalue.OldValueMiddleware()) // 记忆上次表单提交的内容，消费即消失

	// ---------------------------------- 注册路由 ----------------------------------
	// 404
	g.NoRoute(func(c *gin.Context) {
		// controllers.Render404(c)
	})

	r := &router.MyRoute{Router: g}
	// web
	registerWeb(r)
	// api
	registerApi(r)

	return g
}
