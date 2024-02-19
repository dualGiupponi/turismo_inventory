package internal

import (
	"time"

	"github.com/shopspring/decimal"
)

type PromocionPorcentual struct {
	Id                  PromocionId
	Nombre              string
	ListadoAtracciones  []*Atraccion
	PorcentajeDescuento decimal.Decimal
	FechaAlta           time.Time
	FechaActualizacion  time.Time
	FechaBaja           *time.Time
}

func (pp *PromocionPorcentual) GenerarPrecioPromocion() decimal.Decimal {
	var precioTotal decimal.Decimal

	for _, atr := range pp.ListadoAtracciones {
		precioTotal.Add(atr.Costo)
	}

	return precioTotal.Mul(pp.PorcentajeDescuento)
}

func (pp *PromocionPorcentual) ObtenerListado() []*Atraccion {
	return pp.ListadoAtracciones
}

type PromocionPorcentualStorage interface {
	Obtener(PromocionId) (*PromocionPorcentual, error)
	Crear(*PromocionPorcentual) error
	Actualizar(*PromocionPorcentual) error
}
