package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// Restaurant ` vs ' khac nhau the nao vay??
// bien trong golang phai viet HOA de o ngoai moi thay được
type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func main() {
	//fmt.Println("hello")

	dsn := os.Getenv("MYSQL_CONN_STRING")
	//dsn:="food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8&parseTime=True&loc=Local" // data source name
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	//
	//log.Println(db, err)
	//
	fmt.Println(dsn)

	newRestaurant := Restaurant{Name: "Haisan", Addr: "9 Le Loi"}
	if err := db.Create(&newRestaurant).Error; err != nil {
		log.Println(err)
	}
	log.Println("New id:", newRestaurant.Id)

	//read data
	var myRestaurant Restaurant
	if err := db.Where("id=?", 2).First(&myRestaurant).Error; err != nil {
		log.Println(err)
	}
	log.Println(myRestaurant)

	// update
	myRestaurant.Name = "200Labs"
	if err := db.Where("id=?", 2).Updates(&myRestaurant).Error; err != nil {
		log.Println(err)
	}
	log.Println(myRestaurant)

	// delete
	if err := db.Table(Restaurant{}.TableName()).Where("id=?", 1).Delete(nil).Error; err != nil {
		log.Println(err)
	}
	log.Println(myRestaurant)
}
