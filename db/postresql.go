package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectPostgresDb(url string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.TODO(), url)
	if err != nil {
		return nil, err
	}
	err = db.Ping(context.TODO())
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to PostgresSQL!")
	return db, nil
}
