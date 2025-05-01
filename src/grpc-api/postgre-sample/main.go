package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var dbClient *gorm.DB

func init() {
	db, err := gorm.Open(postgres.Open("host=localhost user=user password=password dbname=mydb port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbClient = db

	//migrate()
}

func main() {
	dbClient.Begin()
	//createUser()
	//updateUser()
	//getUsers()
	//getUsersByArray()
	windowFunc()
	dbClient.Commit()
}

func migrate() {
	dbClient.AutoMigrate(&User{}, &Post{})
	err := dbClient.Exec("CREATE TABLE IF NOT EXISTS member_ships (member_ship_id varchar(100) PRIMARY KEY) INHERITS (users)").Error

	if err != nil {
		panic(err)
	}

}

func createUser() {

	profile := Profile{
		Age:   30,
		Hobby: []string{"coding", "gaming"},
	}
	jsonData, _ := json.Marshal(profile)

	user := User{
		UserId:  generateRandomString(20),
		Name:    "test_u" + generateRandomString(5),
		Email:   generateRandomString(5) + "@example.com",
		Profile: datatypes.JSON(jsonData),
	}

	r := dbClient.Create(&user)
	if r.Error != nil {
		fmt.Println("create error", r.Error)
		dbClient.Rollback()
		return
	}

}

func getUsers() {
	var users []User
	r := dbClient.Where("profile ->> 'age' >= ?", "20").Find(&users)
	if r.Error != nil {
		fmt.Println("select error", r.Error)
		dbClient.Rollback()
		return
	}

	for _, v := range users {
		fmt.Println("userData:", v)
	}
}

func updateUser() {
	ids := []string{"6ggYn8o3ScgpTIVlsidu", "dITzZVkV1Np29zbuqOlV", "MsdyqP71xz37aZbb6CHI", "nVjxUmFIhs7qhcUfVJJD"}

	for i, v := range ids {
		updatedProfile := Profile{
			Age:   30 + i + 1,
			Hobby: []string{"Reading", "Swimming"},
		}

		updatedProfileJSON, err := json.Marshal(updatedProfile)
		if err != nil {
			log.Fatal("Failed to marshal profile:", err)
			dbClient.Rollback()
		}

		upd := map[string]any{
			"profile": datatypes.JSON(updatedProfileJSON),
		}

		err = dbClient.Table("users").Where("user_id = ?", v).Updates(upd).Error
		if err != nil {
			log.Fatal("Failed to update profile:", err)
		}
	}

}

func getUsersByArray() {
	var users []User

	err := dbClient.Model(&User{}).
		Where(`profile -> 'hobby' @> ?`, `["coding"]`).
		Find(&users).Error

	if err != nil {
		log.Fatal("Failed to update profile:", err)
	}

	for _, v := range users {
		fmt.Println("userData:", v)
	}
}

func windowFunc() {
	type userWithAvgAge struct {
		UserID string
		RowNum float64
		Age    int
	}

	var users []userWithAvgAge

	dbClient.Raw(`
		SELECT user_id, (profile ->> 'age')::int as age,
			   ROW_NUMBER() OVER (ORDER BY (profile ->> 'age')::int) AS row_num
		FROM users
	`).Scan(&users)

	for _, v := range users {
		fmt.Println("userData:", v)
	}
}

func generateRandomString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
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
