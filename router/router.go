package router

import (
	"fmt"

	ctl "WBABEProject-11/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ct *ctl.Controller) (*Router, error) {
	r := &Router{
		ct: ct,
	}

	return r, nil
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// 실제 인증기능 추가 예정
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}
		auth := c.GetHeader("Authorization")
		fmt.Println("Authorization-word ", auth)

		c.Next()
	}
}

func (p *Router) Index() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	r.GET("/health")

	menu := r.Group("/menu", liteAuth())
	{
		menu.POST("/", p.ct.NewMenuInsert)
		menu.PUT("/", p.ct.UpdateMenu)
		menu.DELETE("/", p.ct.DeleteMenu)
		menu.GET("/", p.ct.GetMenu)
	}
	
	menuReview := r.Group("/menu/review", liteAuth()) 
	{
		menuReview.GET("/", p.ct.GetReview)
		menuReview.POST("/", p.ct.CreateReview)
	}

	order := r.Group("/order", liteAuth())
	{
		order.POST("/", p.ct.CreateOrder)
		order.PUT("/", p.ct.UpdateOrder)
		order.GET("/", p.ct.GetOrder)
		order.GET("/status", p.ct.GetOrderStatus)
	}

	return r
}
