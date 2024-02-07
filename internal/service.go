package internal

import (
	"context"

	"github.com/teohen/rinha-de-backend-2024-q1/internal/domain"
)

type Service interface {
	Create(ctx context.Context, transacao domain.Transacao) (int64, error)
	Get(ctx context.Context, id int64) (domain.Transacao, error)
}

type transacaoService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &transacaoService{
		repo: r,
	}
}
func (ts *transacaoService) Create(ctx context.Context, transacao domain.Transacao) (int64, error) {

	transacaoId, err := ts.repo.Create(ctx, transacao)

	if err != nil {
		return int64(0), err
	}

	return transacaoId, nil
}

func (ts *transacaoService) Get(ctx context.Context, id int64) (domain.Transacao, error) {

	transacao, err := ts.repo.Get(ctx, id)

	if err != nil {
		return domain.Transacao{}, err
	}

	return transacao, nil
}
