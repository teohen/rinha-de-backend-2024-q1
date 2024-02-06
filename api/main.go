package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
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

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = r.Run()

	if err != nil {
		log.Fatal("Unable to create server")
	}

}
