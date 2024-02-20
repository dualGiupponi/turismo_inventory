package internal

import (
	"time"

	"github.com/shopspring/decimal"
)

type ClienteId int
type Cliente struct {
	Id                 ClienteId
	Username           string
	Dinero             decimal.Decimal
	TiempoDisponible   time.Duration
	AtraccionPreferida *TipoAtraccion
	FechaAlta          time.Time
	FechaActualizacion time.Time
	FechaBaja          *time.Time
}

type ClienteStorage interface {
	Obtener(idCliente ClienteId) (*Cliente, error)
	Actualizar(*TipoAtraccion) error
}

type CrearClienteService interface {
	Crear(cliente *Cliente, pwd string) error
}

type ValidarClienteService interface {
	Validar(username, pwd string) error
}

type ActualizarPasswordClienteService interface {
	ActualizarPassword(pwdAnterior, pwdNueva, pwdNuevaConfirmacion string) error
}
