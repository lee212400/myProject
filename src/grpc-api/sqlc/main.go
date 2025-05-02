package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/lee212400/myProject/sqlc/sample"
)

var conn *pgx.Conn

func init() {
	ctx := context.Background()
	c, err := pgx.Connect(ctx, "host=localhost user=user password=password dbname=mydb port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}

	conn = c
}

func main() {
	ctx := context.Background()
	//createSample(ctx)
	getSampe(ctx)
}

func createSample(ctx context.Context) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	queries := sample.New(tx)
	defer tx.Rollback(ctx)

	// create an samples
	sampleDt, err := queries.CreateSample(ctx, sample.CreateSampleParams{
		SampleTitle: "SampleTitle1",
		SampleMemo:  pgtype.Text{String: "SQLC SampleTest", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(sampleDt)

	return tx.Commit(ctx)
}

func getSampe(ctx context.Context) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	queries := sample.New(tx)
	defer tx.Rollback(ctx)

	// create an samples
	sampleDt, err := queries.GetSample(ctx, 1)
	if err != nil {
		return err
	}
	log.Println(sampleDt)

	return tx.Commit(ctx)
}
