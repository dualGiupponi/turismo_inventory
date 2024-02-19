package internal

import "time"

type Compra struct {
	Id                      int
	Comprador               Cliente
	AtraccionesParticulares []*Atraccion
	Promociones             []Promocion
	FechaCompra             time.Time
	FechaBaja               *time.Time
}
