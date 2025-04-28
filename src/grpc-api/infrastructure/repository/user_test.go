package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lee212400/myProject/domain/entity"
	md "github.com/lee212400/myProject/mock/db"
	uc "github.com/lee212400/myProject/utils/context"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryImpl_GetUser(t *testing.T) {
	ct := uc.NewContext(context.Background())
	mockDb := md.NewMockDb()

	type args struct {
		ctx    *entity.Context
		userId string
	}
	tests := []struct {
		name    string
		i       *UserRepositoryImpl
		args    args
		process func()
		want    *entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id"},
			func() {
				mockDb.Mock.ExpectBegin()

				mockDb.Mock.ExpectQuery("select user_id,first_name,last_name,email,age from users where user_id = ?").
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "first_name", "last_name", "email", "age"}).
						AddRow("dummy_user_id", "dummy_first_name", "dummy_last_name", "test@test.com", 30))
			},
			&entity.User{
				UserId:    "dummy_user_id",
				FirstName: "dummy_first_name",
				LastName:  "dummy_last_name",
				Email:     "test@test.com",
				Age:       30,
			}, false,
		},
		{
			"fail: begin error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id"},
			func() {
				mockDb.Mock.ExpectBegin().WillReturnError(errors.New("failed to begin transaction"))
			},
			&entity.User{}, true,
		},
		{
			"fail: not found user",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id"},
			func() {
				mockDb.Mock.ExpectBegin()

				mockDb.Mock.ExpectQuery("select user_id,first_name,last_name,email,age from users where user_id = ?").
					WillReturnRows(sqlmock.NewRows([]string{"user_id", "first_name", "last_name", "email", "age"}))
			},
			&entity.User{}, true,
		},
		{
			"fail: db error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id"},
			func() {
				mockDb.Mock.ExpectBegin()
				mockDb.SelectError(mockDb.Mock)
			},
			&entity.User{}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb.Mock.MatchExpectationsInOrder(false)

			tt.process()

			got, err := tt.i.GetUser(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepositoryImpl.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				require.Equal(t, tt.want.UserId, got.UserId)
				require.Equal(t, tt.want.FirstName, got.FirstName)
				require.Equal(t, tt.want.LastName, got.LastName)
				require.Equal(t, tt.want.Email, got.Email)
				require.Equal(t, tt.want.Age, got.Age)
			}
		})
	}
}

func TestUserRepositoryImpl_CreateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())
	mockDb := md.NewMockDb()
	type args struct {
		ctx       *entity.Context
		firstName string
		lastName  string
		email     string
		age       int32
	}
	tests := []struct {
		name    string
		i       *UserRepositoryImpl
		args    args
		process func()
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_first_name", "dummy_last_name", "test@test.com", 30},
			func() {
				mockDb.Mock.ExpectBegin()

				mockDb.Mock.ExpectExec("INSERT INTO users (user_id, first_name, last_name, email, age) VALUES (?, ?, ?, ?, ?)").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			"dummy_user_id", false,
		},
		{
			"fail: begin error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_first_name", "dummy_last_name", "test@test.com", 30},
			func() {
				mockDb.Mock.ExpectBegin().WillReturnError(errors.New("failed to begin transaction"))
			},
			"", true,
		},
		{
			"fail: db error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_first_name", "dummy_last_name", "test@test.com", 30},
			func() {
				mockDb.Mock.ExpectBegin()
				mockDb.CreateError(mockDb.Mock)
			},
			"", true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb.Mock.MatchExpectationsInOrder(false)
			tt.process()

			_, err := tt.i.CreateUser(tt.args.ctx, tt.args.firstName, tt.args.lastName, tt.args.email, tt.args.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepositoryImpl.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepositoryImpl_UpdateUser(t *testing.T) {
	ct := uc.NewContext(context.Background())
	mockDb := md.NewMockDb()

	type args struct {
		ctx    *entity.Context
		userId string
		age    int32
	}
	tests := []struct {
		name    string
		i       *UserRepositoryImpl
		args    args
		process func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id", 30},
			func() {
				mockDb.Mock.ExpectBegin()

				mockDb.Mock.ExpectExec("update users set age = ? where user_id = ?").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			false,
		},
		{
			"fail: begin error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id", 30},
			func() {
				mockDb.Mock.ExpectBegin().WillReturnError(errors.New("failed to begin transaction"))
			},
			true,
		},
		{
			"fail: db error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id", 30},
			func() {
				mockDb.Mock.ExpectBegin()
				mockDb.UpdateError(mockDb.Mock)
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb.Mock.MatchExpectationsInOrder(false)
			tt.process()

			if err := tt.i.UpdateUser(tt.args.ctx, tt.args.userId, tt.args.age); (err != nil) != tt.wantErr {
				t.Errorf("UserRepositoryImpl.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepositoryImpl_DeleteUser(t *testing.T) {
	ct := uc.NewContext(context.Background())
	mockDb := md.NewMockDb()
	type args struct {
		ctx    *entity.Context
		userId string
	}
	tests := []struct {
		name    string
		i       *UserRepositoryImpl
		args    args
		process func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"correct",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id"},
			func() {
				mockDb.Mock.ExpectBegin()

				mockDb.Mock.ExpectExec("delete from users where user_id = ?").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			false,
		},
		{
			"fail: begin error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id"},
			func() {
				mockDb.Mock.ExpectBegin().WillReturnError(errors.New("failed to begin transaction"))
			},
			true,
		},
		{
			"fail: db error",
			NewUserRepositoryImpl(mockDb),
			args{ct, "dummy_user_id"},
			func() {
				mockDb.Mock.ExpectBegin()
				mockDb.DeleteError(mockDb.Mock)
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDb.Mock.MatchExpectationsInOrder(false)
			tt.process()

			if err := tt.i.DeleteUser(tt.args.ctx, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("UserRepositoryImpl.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
