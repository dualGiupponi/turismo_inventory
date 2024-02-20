package postgre

import (
	"database/sql"
	"github.com/dualgiupponi/turismo_inventory/internal"
)

const (
	ObtenerTipoAtraccionSql = "SELECT * FROM tipos_atracciones WHERE tip_atr_id = @tip_atr_id"
)

type tipoAtraccionDb struct {
	tip_atr_id      int16
	tip_atr_nom     string
	tip_atr_fec_alt sql.NullTime
	tip_atr_fec_baj sql.NullTime
}

func tipoAtraccionDbToTipoAtraccion(taDb tipoAtraccionDb) internal.TipoAtraccion {
	ta := internal.TipoAtraccion{
		Id:        internal.TipoAtraccionId(taDb.tip_atr_id),
		Nombre:    taDb.tip_atr_nom,
		FechaAlta: taDb.tip_atr_fec_alt.Time,
	}

	if taDb.tip_atr_fec_baj.Valid {
		ta.FechaBaja = &taDb.tip_atr_fec_baj.Time
	}

	return ta
}

func tipoAtraccionToTipoAtraccionDb(ta internal.TipoAtraccion) tipoAtraccionDb {
	return tipoAtraccionDb{
		tip_atr_id:      int16(ta.Id),
		tip_atr_nom:     ta.Nombre,
		tip_atr_fec_alt: sql.NullTime{Time: ta.FechaAlta},
		tip_atr_fec_baj: sql.NullTime{Time: *ta.FechaBaja},
	}
}
