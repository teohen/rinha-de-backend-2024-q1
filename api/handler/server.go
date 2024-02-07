package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teohen/rinha-de-backend-2024-q1/internal"
)

type ServerAPI struct {
	Server *http.Server
}

func NewServer(conn *pgxpool.Pool) {
	port := os.Getenv("HTTP_PORT")

	r := gin.Default()

	repository := internal.NewTransacaoRepo(conn)

	service := internal.NewService(repository)

	handler := NewTransacaoHandler(service)

	r.POST("/clientes/:id/transacoes", handler.Create)

	err := r.Run(port)

	if err != nil {
		log.Fatal("Unable to create server")
	}

}
