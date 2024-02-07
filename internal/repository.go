package internal

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teohen/rinha-de-backend-2024-q1/internal/domain"
)

type Repository interface {
	Get(ctx context.Context, id int64) (domain.Transacao, error)
	Create(ctx context.Context, transacao domain.Transacao) (int64, error)
}

type transacaoRepository struct {
	db *pgxpool.Pool
}

func NewTransacaoRepo(db *pgxpool.Pool) Repository {
	return &transacaoRepository{
		db: db,
	}
}

func (tr *transacaoRepository) Get(ctx context.Context, id int64) (domain.Transacao, error) {
	transacao := domain.Transacao{}

	get := "SELECT  id, client_id, valor, tipo, descricao, realizada_em FROM transacoes where = $1"

	row := tr.db.QueryRow(ctx, get, id)

	err := row.Scan(&transacao.IdCliente, &transacao.Valor, &transacao.Tipo, &transacao.Descricao, &transacao.RealizadaEm)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return domain.Transacao{}, nil
		}

		if err != nil {
			return transacao, fmt.Errorf("get transacao: %w", err)
		}
	}

	return transacao, nil
}

func (tr *transacaoRepository) Create(ctx context.Context, transacao domain.Transacao) (int64, error) {

	insert := "INSERT INTO transacao (client_id, valor, tipo, descricao) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING returning id;"

	err := tr.db.QueryRow(ctx, insert, transacao.IdCliente, transacao.Valor, transacao.Tipo, transacao.Descricao).Scan(transacao.Id)

	if err != nil {
		return 0, fmt.Errorf("create transacao: %w", err)
	}

	return transacao.Id, nil
}
