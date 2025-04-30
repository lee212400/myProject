package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func init() {
	db, err := gorm.Open(postgres.Open("host=localhost user=user password=password dbname=mydb port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbClient = db

	migrate()
}

func main() {
	createUser()
}

func migrate() {
	dbClient.AutoMigrate(&User{}, &Post{})
}

func createUser() {
	dbClient.Begin()
	profile := Profile{
		Age:   30,
		Hobby: []string{"coding", "gaming"},
	}
	jsonData, _ := json.Marshal(profile)

	user := User{
		UserId:  "user_2",
		Name:    "test_user",
		Email:   "test@example.com",
		Profile: datatypes.JSON(jsonData),
	}

	r := dbClient.Create(&user)
	if r.Error != nil {
		dbClient.Rollback()
		return
	}

	dbClient.Commit()

}

func getUsers() {
	dbClient.Begin()
	var users []User
	r := dbClient.Where("profile ->> 'age' = ?", "30").Find(&users)
	if r.Error != nil {
		dbClient.Rollback()
		return
	}
	dbClient.Commit()

	for _, v := range users {
		fmt.Println("userData:", v)
	}
}

type Profile struct {
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

type User struct {
	UserId  string         `gorm:"type:varchar(100);primaryKey"`
	Name    string         `gorm:"type:varchar(100)"`
	Email   string         `gorm:"type:varchar(100)"`
	Profile datatypes.JSON `gorm:"type:jsonb"` // GORMのjsonb type
}

type Post struct {
	PostId  string `gorm:"type:varchar(100);primaryKey"`
	UserId  string `gorm:"type:varchar(100);primaryKey"`
	Title   string `gorm:"type:varchar(100)"`
	Content string `gorm:"type:text"`
}
