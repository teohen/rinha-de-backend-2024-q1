package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/teohen/rinha-de-backend-2024-q1/api/handler"
)

func main() {

	godotenv.Load()

	var psqlconn string = fmt.Sprintf("host=%, user=% password=% dbname=%", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	poolConfig, err := pgxpool.ParseConfig(psqlconn)

	if err != nil {
		log.Fatal("Couldnt parse config", err)
	}

	poolConfig.MaxConns = 100
	poolConfig.MinConns = 10

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)

	if err != nil {
		log.Fatal("couldnt create conn pool", err)
	}

	defer db.Close()

	handler.NewServer(db)
}
