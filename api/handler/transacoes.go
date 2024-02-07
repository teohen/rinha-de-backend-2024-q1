package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teohen/rinha-de-backend-2024-q1/internal"
	"github.com/teohen/rinha-de-backend-2024-q1/internal/domain"
)

type TransacaoRequest struct {
	Valor     int64  `json:"valor"`
	Tipo      rune   `json:"tipo"`
	Descricao string `json:"descricao"`
}

type Handler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
}

type transacaoHandler struct {
	service internal.Service
}

func NewTransacaoHandler(transacaoService internal.Service) Handler {
	return &transacaoHandler{
		service: transacaoService,
	}
}

func (th *transacaoHandler) Create(c *gin.Context) {

	paramClienteId := c.Param("id")
	clienteId, err := strconv.Atoi(paramClienteId)

	if err != nil {
		c.String(http.StatusBadRequest, "ERRO CARAI")
		return
	}

	var transacaoParams TransacaoRequest

	if err := c.BindJSON(&transacaoParams); err != nil {
		log.Fatal("FFALHA NO MARSAL")
		return
	}

	if err != nil {
		c.String(http.StatusBadRequest, "ERRO CARAI")
		return
	}

	newTransacao := domain.Transacao{
		Tipo:      transacaoParams.Tipo,
		IdCliente: clienteId,
		Descricao: transacaoParams.Descricao,
		Valor:     transacaoParams.Valor,
	}

	transacaoId, err := th.service.Create(context.Background(), newTransacao)

	fmt.Println("ID DA TRANSACAO", transacaoId)

	if err != nil {
		c.String(http.StatusBadRequest, "ERRO CARAI")
		return
	}

	c.String(http.StatusOK, "DEUC ERTO")
}

func (th *transacaoHandler) Get(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 0)

	if err != nil {
		c.String(http.StatusBadRequest, "ERRO CARAI")
		return
	}

	transacao, _ := th.service.Get(context.Background(), id)
	fmt.Println("TESTE", transacao)

	c.String(http.StatusOK, "DEU CETO")
}
