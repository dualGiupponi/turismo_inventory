package internal

import (
	"time"

	"github.com/shopspring/decimal"
)

type PromocionPrecioFijo struct {
	Id                 PromocionId
	Nombre             string
	ListadoAtracciones []*Atraccion
	MontoTotal         decimal.Decimal
	FechaAlta          time.Time
	FechaActualizacion time.Time
	FechaBaja          *time.Time
}

func (ppp *PromocionPrecioFijo) GenerarPrecioPromocion() decimal.Decimal {
	return ppp.MontoTotal
}

func (ppp *PromocionPrecioFijo) ObtenerListado() []*Atraccion {
	return ppp.ListadoAtracciones
}

type PromocionPrecioFijoStorage interface {
	Obtener(PromocionId) (*PromocionPrecioFijo, error)
	Crear(*PromocionPrecioFijo) error
	Actualizar(*PromocionPrecioFijo) error
}
