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
	sName := c.PostForm("name")
	sPrice := c.PostForm("price")

	if len(sName) <= 0 {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", nil)
		return
	}

	menu, _ := p.md.GetOneMenu("name", sName) 
	fmt.Println("res ", menu)
	if menu == (model.Menu{}) {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "could not found person", nil)
		return
	}

	nPrice, err := strconv.Atoi(sPrice)
	if err != nil {
		nPrice = 1
	}

	if err := p.md.UpdateMenu(sName, nPrice); err != nil {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
	c.Next()
}

func (p *Controller) DeleteMenu(c *gin.Context) {
	sName := c.Param("name")

	if len(sName) <= 0 {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", nil)
		return
	}

	_, err := p.md.GetOneMenu("name", sName)
	if err != nil {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "exist resistery person", nil)
		return
	}

	if err := p.md.DeleteMenu(sName); err != nil {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "fail delete db", err)
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"result": "ok",
	})
	c.Next()
}

func (p *Controller) GetMenu(c *gin.Context) {
	if menu, err := p.md.GetMenu(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res":  "fail",
			"body": err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"res":  "ok",
			"body": menu,
		})
		c.Next()
	}
}

func (p *Controller) GetOneMenu(c *gin.Context) {
	sName := c.Param("name")
	if len(sName) <= 0 {
		p.RespError(c, nil, 400, "fail, Not Found Param", nil)
		c.Abort()
		return
	}

	if menu, err := p.md.GetOneMenu("name", sName); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res":  "fail",
			"body": err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"res":  "ok",
			"body": menu,
		})
		c.Next()
	}
}


func (p *Controller) GetReview(c *gin.Context) {
	sName := c.Param("name")
	if len(sName) <= 0 {
		p.RespError(c, nil, 400, "fail, Not Found Param", nil)
		c.Abort()
		return
	}

	if review, err := p.md.GetReview("name", sName); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res":  "fail",
			"body": err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"res":  "ok",
			"body": review,
		})
		c.Next()
	}

}

func (p *Controller) NewReviewInsert(c *gin.Context) {
	name := c.PostForm("name")
	menu := c.PostForm("menu")
	rating := c.PostForm("rating")
	orderNumber := c.PostForm("ordernumber")
	review := c.PostForm("review")

	if len(name) <= 0 || len(menu) <= 0 {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", nil)
		return
	}

	nrating, err := strconv.Atoi(rating)
	if err != nil {
		nrating = 1
	}

	nordernumber, err := strconv.Atoi(orderNumber)
	if err != nil {
		nordernumber = 1
	}

	req := model.Review{Name: name, Menu: menu, Rating: nrating, OrderNumber: nordernumber, Review: review}
	if err := p.md.CreateReview(req); err != nil {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
	c.Next()
}


func (p *Controller) NewOrderInsert(c *gin.Context) {
	name := c.PostForm("name")
	menu := c.PostForm("menu")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	status := c.PostForm("status")

	if len(name) <= 0 || len(menu) <= 0 {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", nil)
		return
	}

	nstatus, err := strconv.Atoi(status)
	if err != nil {
		nstatus = 1
	}

	req := model.Order{Menu: menu, Name: name, Phone: phone, Address: address, Status: nstatus}
	if err := p.md.CreateOrder(req); err != nil {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
	c.Next()

}

func (p *Controller) UpdateOrder(c *gin.Context) {
	sName := c.PostForm("name")
	sMenu := c.PostForm("menu")

	if len(sName) <= 0 {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", nil)
		return
	}

	order, _ := p.md.GetOrder("name", sName) 
	fmt.Println("res ", order)
	if order == (model.Order{}) {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "could not found person", nil)
		return
	}

	if err := p.md.UpdateOrder(sName, sMenu); err != nil {
		p.RespError(c, nil, http.StatusUnprocessableEntity, "parameter not found", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
	c.Next()
	
}

func (p *Controller) GetOrder(c *gin.Context) {
	sName := c.Param("name")
	if len(sName) <= 0 {
		p.RespError(c, nil, 400, "fail, Not Found Param", nil)
		c.Abort()
		return
	}

	if order, err := p.md.GetOrder("name", sName); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res":  "fail",
			"body": err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"res":  "ok",
			"body": order,
		})
		c.Next()
	}

}

func (p *Controller) GetOrderStatus(c *gin.Context) {

}
