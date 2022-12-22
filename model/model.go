package model

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client *mongo.Client
	colMenu *mongo.Collection
	colOrder *mongo.Collection
}

// 이것을 이렇게 구조를 가져가도 되나?
type Menu struct {
	Name string `json:"name" bson:"name"`
	Soldout int `json:"soldout" bson:"soldout"`
	Stock int `json:"stock" bson:"stock"`
	Origin string `json:"origin" bson:"origin"`
	Price int `json:"price" bson:"price"`
	Rating int `json:"rating" bson:"rating"`
	OrderNumber int `json:"orderNumber" bson:"ordernumber"`
	ReorderNumber int `json:"reordernumber" bson:"reordernumber"`
	Review string `json:"review" bson:"review"`
}

type Order struct {
	Menu string `json:"menu" bson:"menu"`
	Name string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
	Address	string `json:"address" bson:"address"`
	Status int `json:"status" bson:"status"`
}

func NewModel() (*Model, error) {
	r := &Model{}

	mgUrl := "mongodb://127.0.0.1:27017"

	var err error
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if  err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-ready") 
		r.colMenu = db.Collection("tMenu")
		r.colOrder = db.Collection("tOrder")
	}
	return r, nil
}

func (p *Model) CreateMenu(menu Menu) error {
	if _, err := p.colMenu.InsertOne(context.TODO(), menu); err != nil {
		fmt.Println("fail insert new menu")
		return fmt.Errorf("fail, insert new menu")
	}
	return nil
}

func (p *Model) UpdateMenu(name string, price int) error {
	filter := bson.M{"name": name}
	update := bson.M{
		"$set": bson.M{
			"price": price, 
		},
	}

	if _, err := p.colMenu.UpdateOne(context.Background(), filter, update); err != nil {
		return err
	}
	return nil
}

func (p *Model) DeleteMenu(name string) error {
	filter := bson.M{"name": name}

	if res, err := p.colMenu.DeleteOne(context.TODO(), filter); res.DeletedCount <= 0 {
		return fmt.Errorf("Could not Delete, Not found name %s", name)
	} else if err != nil {
		return err
	}
	return nil
}

// 추가 조사 필요
func (p *Model) GetMenu() {
	filter := bson.D{}
	cursor, err := p.colMenu.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var menu []Menu
	for _, result := range menu {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}

// 메뉴 하나 불러오기
func (p *Model) GetOneMenu(flag, elem string) (Menu, error) {
	opts := []*options.FindOneOptions{}
	
	filter := bson.M{}
	if flag == "name" {
		filter = bson.M{"name": elem}
	}

	var menu Menu
	if err := p.colMenu.FindOne(context.TODO(), filter, opts...).Decode(&menu); err != nil {
		return menu, err
	} else {
		return menu, err
	}
}

// 리뷰만 불러올 수 있도록 세팅 필요
func (p *Model) GetReview() {
	filter := bson.D{}
	cursor, err := p.colMenu.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var menu []Menu
	for _, result := range menu {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}

func (p *Model) CreateReview(menu Menu) error {
	if _, err := p.colMenu.InsertOne(context.TODO(), menu); err != nil {
		fmt.Println("fail insert new review")
		return fmt.Errorf("fail, insert new review")
	}
	return nil
}

func (p *Model) CreateOrder(order Order) error {
	if _, err := p.colMenu.InsertOne(context.TODO(), order); err != nil {
		fmt.Println("fail insert new order")
		return fmt.Errorf("fail, insert new order")
	}
	return nil
}

func (p *Model) UpdateOrder(name, menu string) error {
	filter := bson.M{"name": name}
	update := bson.M{
		"$set": bson.M{
			"menu": menu, 
		},
	}

	if _, err := p.colOrder.UpdateOne(context.Background(), filter, update); err != nil {
		return err
	}
	return nil
}

func (p *Model) GetOrder() {
	filter := bson.D{}
	cursor, err := p.colOrder.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var order []Order
	for _, result := range order {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

}

func (p *Model) GetOrderStatus(name string) (Order, error) {
	filter := bson.M{"name": name}

	var order Order
	if err := p.colOrder.FindOne(context.TODO(), filter).Decode(&order); err != nil {
		return order, err
	} else {
		return order, err
	}
}
