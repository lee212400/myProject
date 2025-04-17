package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/infrastructure/db"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (i *UserRepositoryImpl) GetUser(ctx *entity.Context, userId string) (*entity.User, error) {
	fmt.Println("Repository GetUser")
	db, err := db.GetDb(ctx, i.db)
	if err != nil {
		return &entity.User{}, err
	}

	query := "select user_id,first_name,last_name,email,age from users where user_id = ?"

	r := db.QueryRowContext(ctx.Ctx, query, userId)
	var user entity.User
	err = r.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return &entity.User{}, fmt.Errorf("user not found")
		}
		return &entity.User{}, err
	}

	return &user, nil
}
func (i *UserRepositoryImpl) CreateUser(ctx *entity.Context, firstName string, lastName string, email string, age int32) (string, error) {
	db, err := db.GetDb(ctx, i.db)
	if err != nil {
		return "", err
	}

	query := `
		INSERT INTO users (user_id, first_name, last_name, email, age)
		VALUES (?, ?, ?, ?, ?)
	`
	userId := generateRandomString(20)

	_, err = db.ExecContext(ctx.Ctx, query, userId, firstName, lastName, email, age)
	if err != nil {
		return "", err
	}

	return userId, nil
}
func (i *UserRepositoryImpl) UpdateUser(ctx *entity.Context, userId string, age int32) error {
	db, err := db.GetDb(ctx, i.db)
	if err != nil {
		return err
	}

	query := `update users set age = ? where user_id = ?`

	_, err = db.Exec(query, age, userId)

	return err
}
func (i *UserRepositoryImpl) DeleteUser(ctx *entity.Context, userId string) error {
	db, err := db.GetDb(ctx, i.db)
	if err != nil {
		return err
	}

	query := `delete from users where user_id = ?`

	_, err = db.Exec(query, userId)

	return err
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
