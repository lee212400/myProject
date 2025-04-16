package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/lee212400/myProject/rpc"
)

type server struct {
	pb.UnimplementedSampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	tx, err := connectDB()
	if err != nil {
		return &pb.HelloResponse{}, err
	}

	u, err := getUser(tx)
	if err != nil {
		return &pb.HelloResponse{}, err
	}

	fmt.Println("get user data", u)

	return &pb.HelloResponse{
		Message: u[0],
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterSampleServiceServer(s, &server{})
	log.Println("Server running at :50051")

	reflection.Register(s)

	s.Serve(lis)
}

func init() {
	var err error
	tx, err := connectDB()
	defer recoverFromPanic()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err != nil {
		panic(err)
	}
	err = createTb(tx)
	if err != nil {
		panic(err)
	}
	err = insertData(tx)
	if err != nil {
		panic(err)
	}

}

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("error:%v\n", err)
		fmt.Printf("retry... count:%v\n", count)

		return open(path, count)
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() (*sql.Tx, error) {
	var path string = fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=true",
		"root", "password", "sampledb")

	return open(path, 100).Begin()
}

func createTb(tx *sql.Tx) error {
	var tableName string
	checkQuery := `
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = DATABASE() AND table_name = 'users';
	`

	err := tx.QueryRow(checkQuery).Scan(&tableName)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check table existence: %w", err)
	}

	if tableName != "" {
		return nil
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL
	);
	`

	_, err = tx.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return err
}

func insertData(tx *sql.Tx) error {
	insertQuery := `
	INSERT INTO users (name, email) VALUES 
		('testuser1', 'test@example.com'),
		('testuser2', 'test2@example.com');
	`

	_, err := tx.Exec(insertQuery)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return err
}

func getUser(tx *sql.Tx) ([]string, error) {
	u := []string{}
	query := `SELECT name FROM users`

	rows, err := tx.Query(query)
	if err != nil {
		return u, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return u, err
		}
		u = append(u, name)
	}

	if err := rows.Err(); err != nil {
		return u, err
	}

	return u, nil
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
