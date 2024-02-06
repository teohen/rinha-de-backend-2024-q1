package domain

import "github.com/google/uuid"

type Transacao struct {
	Id        uuid.UUID `json:"id"`
	Valor     int64     `json:"valor"`
	Descricao string    `json:"descricao"`
}
