package routers

import (
	"blog/middleware"
	"blog/utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "static/admin/index.html")
	p.AddFromFiles("front", "static/front/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.HTMLRender = createMyRender()

	r.Static("/static", "./static/front/static")
	r.Static("/admin", "./static/admin")
	r.StaticFile("/favicon.ico", "static/front/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	_ = r.Run(utils.HttpPort)

	//auth := r.Group("ap1/v1")
}
