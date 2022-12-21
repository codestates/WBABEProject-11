package model

import (
	"context"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client *mongo.Client
	colMenu *mongo.Collection
	colOrder *mongo.Collection
}

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

func CreateMenu() {

}

func UpdateMenu() {

}

func DeleteMenu() {

}

func GetMenu() {

}

func GetReview() {

}

func CreateReview() {

}

func CreateOrder() {

}

func UpdateOrder() {

}

func GetOrder() {

}

func GetOrderStatus() {

}
