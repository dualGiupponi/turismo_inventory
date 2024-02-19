package internal

import "time"

type TipoAtraccionId int16
type TipoAtraccion struct {
	Id        TipoAtraccionId
	Nombre    string
	FechaAlta time.Time
	FechaBaja *time.Time
}

type TipoAtraccionStorage interface {
	Obtener(idTipoAtraccion int16) (*TipoAtraccion, error)
	ObtenerListado() ([]*TipoAtraccion, error)
	Crear(*TipoAtraccion) error
	Actualizar(*TipoAtraccion) error
}
