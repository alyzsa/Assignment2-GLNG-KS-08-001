package database

import (
	"fmt"
	"log"
	"github.com/alyzsa/Assignment2-GLNG-KS-08-001/models"


	_"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

var (
	db *gorm.DB
	err error 
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "alysha09"
	dbname= "postgres"

)

func StartDB(){
	var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbname)
	
	db, err=gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil{
		log.Fatal("error connecting to database: ", err)
	}
	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	fmt.Println("Connected to Database")

}

func GetDB()*gorm.DB{
	return db
}