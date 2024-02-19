package internal

import "github.com/shopspring/decimal"

type PromocionId int

type Promocion interface {
	Generar() decimal.Decimal
	ObtenerListado() []*Atraccion
}
