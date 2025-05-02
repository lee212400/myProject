package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lee212400/myProject/domain/entity"
	uc "github.com/lee212400/myProject/utils/context"
)

func main() {

	dbClient := NewDB()
	newCtx := uc.NewContext(context.Background())
	err := getSample(newCtx, dbClient)
	//err := createSample(newCtx, dbClient)
	if err != nil {
		log.Fatal("error:", err)
		(*dbClient).Rollback(newCtx)
	}

	(*dbClient).Commit(newCtx)

}

func createSample(ctx *entity.Context, db *pgx.Tx) error {
	s := &sample{}
	err := (*db).QueryRow(ctx,
		"INSERT INTO samples (sample_title, sample_memo) VALUES ($1, $2) RETURNING *",
		"sample_title2", "sampe_memo2",
	).Scan(&s.ID, &s.SampleTitle, &s.SampleMemo)

	fmt.Println("create data: ", s)
	return err
}

func getSample(ctx *entity.Context, db *pgx.Tx) error {
	rows, err := (*db).Query(ctx, "SELECT id, sample_title, sample_memo FROM samples")
	if err != nil {
		return err
	}
	defer rows.Close()

	samples := []*sample{}
	for rows.Next() {
		s := &sample{}
		if err := rows.Scan(&s.ID, &s.SampleTitle, &s.SampleMemo); err != nil {
			return err
		}
		samples = append(samples, s)
	}

	for _, v := range samples {
		fmt.Println("sample data : ", v)
	}
	return nil
}

// クエリを一括で送信する
func bulkCreate(ctx *entity.Context, db *pgx.Tx) error {
	samples := []*sample{
		{
			SampleTitle: "testTitleData",
			SampleMemo:  "testMemoData",
		},
	}
	batch := &pgx.Batch{}
	for _, s := range samples {
		batch.Queue(`INSERT INTO sample (sample_title, sample_memo) VALUES ($1, $2)`, s.SampleTitle, s.SampleMemo)
	}
	br := (*db).SendBatch(ctx, batch)
	err := br.Close()

	return err
}

// データをパケット単位で一括で送信する
// エラーが発生したらすべてのデータが失敗になる
func copyCreate(ctx *entity.Context, db *pgx.Tx) error {

	rows := [][]interface{}{
		{"title 1", "memo 1"},
		{"title 2", "memo 2"},
		{"title 3", "memo 3"},
	}

	inputRows := pgx.CopyFromRows(rows)

	_, err := (*db).CopyFrom(
		ctx,
		pgx.Identifier{"samples"},
		[]string{"sample_title", "sample_memo"},
		inputRows,
	)

	return err
}

func NewDB() *pgx.Tx {
	dsn := "host=localhost user=user password=password dbname=mydb port=5432 sslmode=disable"
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	cfg.MaxConns = 10
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Hour

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pgPool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		panic(err)
	}

	db, err := pgPool.Begin(ctx)
	if err != nil {
		panic(err)
	}
	return &db

}

type sample struct {
	ID          int64
	SampleTitle string
	SampleMemo  string
}
