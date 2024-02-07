package domain

type Transacao struct {
	Id          int64
	IdCliente   int
	Valor       int64
	Descricao   string
	Tipo        rune
	RealizadaEm string
}
