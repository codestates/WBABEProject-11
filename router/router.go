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
		menu.DELETE("/:name", p.ct.DeleteMenu)
		menu.GET("/", p.ct.GetMenu)
		menu.GET("/:name", p.ct.GetOneMenu)
	}
	
	menuReview := r.Group("/menu/review", liteAuth()) 
	{
		menuReview.GET("/:name", p.ct.GetReview)
		menuReview.POST("/", p.ct.NewReviewInsert)
	}

	order := r.Group("/order", liteAuth())
	{
		order.POST("/", p.ct.NewOrderInsert)
		order.PUT("/", p.ct.UpdateOrder)
		order.GET("/:name", p.ct.GetOrder)
		order.GET("/status", p.ct.GetOrderStatus)
	}
	/* [코드리뷰]
	 * Group을 사용하여 API 성격에 따라 request를 관리하는 코드는 매우 좋은 코드입니다.
     * 일반적으로 현업에서도 이와 같은 코드를 자주 사용합니다. 훌륭합니다.
	 *
	 * 코드의 확장성을 고려하였을때, endpoint 관리를 함께 고려한 코드를 개발하는 것도 추천드립니다.
	 * 예를들어 /order/status 를 호출하는 클라이언트(Web, App, etc..)들이 실시간으로 들어오고 있을 때,
	 * controller의 GetOrderStatus function을 변경해야 하는 상황이 발생한다면,
	 * /order/status2 로 받아주는 경우가 있을 것이고(/order/status는 그대로 받아주면서)
	 * 처음부터 /order/v1/status 로 관리되며, /order/v2/status 리뉴얼 버전에 따라 version up을 시켜
	 * v01 방식의 클라이언트와, v02 방식의 클라이언트를 모두 받아줄 수 있는 확장성 있는 코드를 구현해보시는 것을 추천드립니다.
	 */

	return r
}
