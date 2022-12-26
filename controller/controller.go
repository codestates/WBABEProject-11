package controller

import (
	"WBABEProject-11/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{
		md: rep,
	}
	return r, nil
}

func (p *Controller) RespError(c *gin.Context, body interface{}, status int, err ...interface{}) {
	bytes, _ := json.Marshal(body)

	fmt.Println("Request error", "path", c.FullPath(), "body", bytes, "status", status, "error", err)

	c.JSON(status, gin.H{
		"Error": "Request Error",
		"path": c.FullPath(),
		"body": bytes,
		"status": status,
		"error": err,
	})
	c.Abort()
}

// menu와 review 구조를 어떻게 가져갈 것인가?
func (p *Controller) NewMenuInsert(c *gin.Context) {
	name := c.PostForm("name")
	soldout := c.PostForm("soldout")
	stock := c.PostForm("stock")
	origin := c.PostForm("origin")
	price := c.PostForm("price")

	if len(name) <= 0 || len(price) <= 0 {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", nil)
		return
	}

	menu, _ := p.md.GetOneMenu("name", name)
	if menu != (model.Menu{}) {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "already resistery menu", nil)
		return
	}

	nSoldout, err := strconv.Atoi(soldout)
	if err != nil {
		nSoldout = 1
	}

	nstock, err := strconv.Atoi(stock)
	if err != nil {
		nstock = 1
	}

	nprice, err := strconv.Atoi(price)
	if err != nil {
		nprice = 1
	}

	req := model.Menu{Name: name, Soldout: nSoldout, Stock: nstock, Origin: origin, Price: nprice}
	if err := p.md.CreateMenu(req); err != nil {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
	c.Next()
}

func (p *Controller) UpdateMenu(c *gin.Context) {

}

func (p *Controller) DeleteMenu(c *gin.Context) {


}

func (p *Controller) GetMenu(c *gin.Context) {

}

func (p *Controller) GetReview(c *gin.Context) {

}

func (p *Controller) NewReviewInsert(c *gin.Context) {

}

func (p *Controller) NewOrderInsert(c *gin.Context) {

}

func (p *Controller) UpdateOrder(c *gin.Context) {

}

func (p *Controller) GetOrder(c *gin.Context) {

}

func (p *Controller) GetOrderStatus(c *gin.Context) {

}
