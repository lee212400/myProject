package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/infrastructure/db"
	ue "github.com/lee212400/myProject/utils/errors"
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
		return &entity.User{}, ue.WithError(ctx, ue.Internal, err)
	}

	query := "select user_id,first_name,last_name,email,age from users where user_id = ?"

	r := db.QueryRowContext(ctx, query, userId)
	var user entity.User
	err = r.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return &entity.User{}, ue.New(ctx, ue.NotFound, "No found user")
		}
		return &entity.User{}, ue.WithError(ctx, ue.Internal, err)
	}

	return &user, nil
}
func (i *UserRepositoryImpl) CreateUser(ctx *entity.Context, firstName string, lastName string, email string, age int32) (string, error) {
	db, err := db.GetDb(ctx, i.db)
	if err != nil {
		return "", ue.WithError(ctx, ue.Internal, err)
	}

	query := `
		INSERT INTO users (user_id, first_name, last_name, email, age)
		VALUES (?, ?, ?, ?, ?)
	`
	userId := generateRandomString(20)

	_, err = db.ExecContext(ctx, query, userId, firstName, lastName, email, age)
	if err != nil {
		return "", ue.WithError(ctx, ue.Internal, err)
	}

	return userId, nil
}
func (i *UserRepositoryImpl) UpdateUser(ctx *entity.Context, userId string, age int32) error {
	db, err := db.GetDb(ctx, i.db)
	if err != nil {
		return ue.WithError(ctx, ue.Internal, err)
	}

	query := `update users set age = ? where user_id = ?`

	_, err = db.Exec(query, age, userId)
	if err != nil {
		return ue.WithError(ctx, ue.Internal, err)
	}
	return nil
}
func (i *UserRepositoryImpl) DeleteUser(ctx *entity.Context, userId string) error {
	db, err := db.GetDb(ctx, i.db)
	if err != nil {
		return err
	}

	query := `delete from users where user_id = ?`

	_, err = db.Exec(query, userId)
	if err != nil {
		return ue.WithError(ctx, ue.Internal, err)
	}
	return nil
}

func generateRandomString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}
