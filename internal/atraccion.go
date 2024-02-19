package internal

import (
	"time"

	"github.com/shopspring/decimal"
)

type AtraccionId int
type Atraccion struct {
	Id                 AtraccionId
	Nombre             string
	Duracion           time.Duration
	Costo              decimal.Decimal
	CapacidadMaxima    int16
	TipoAtraccion      TipoAtraccion
	FechaAlta          time.Time
	FechaActualizacion time.Time
	FechaBaja          *time.Time
}

type AtraccionStorage interface {
	Obtener(idAtraccion AtraccionId) (*Atraccion, error)
	ObtenerListado() ([]*Atraccion, error)
	Crear(*Atraccion) error
	Actualizar(*Atraccion) error
}
