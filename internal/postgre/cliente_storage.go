package postgre

import (
	"database/sql"
	"errors"
	"github.com/dualgiupponi/turismo_inventory/internal"
	"time"

	"github.com/shopspring/decimal"
)

const (
	ObtenerClienteSql string = `SELECT * FROM clientes 
    								INNER JOIN tipos_atracciones 
    					   				ON cli_tip_atr_id = tip_atr_id 
         				 		WHERE cli_id = @cli_id;`
	CrearClienteSql string = `INSERT INTO clientes (cli_id, cli_usrnm, cli_din_disp, cli_tiempo_disp, cli_tip_atr_id, cli_fec_alt, cli_fec_act, cli_fec_baj)
								VALUES(@cli_id, @cli_usrnm, @cli_din_disp, @cli_tiempo_disp, @cli_tip_atr_id, @cli_fec_alt, @cli_fec_act, @cli_fec_baj);`
)

type ClienteStorage struct {
	*DbConnection
}

func NewClienteStorage(dbConn *DbConnection) *ClienteStorage {
	return &ClienteStorage{DbConnection: dbConn}
}

func (cl *ClienteStorage) Obtener(idCliente internal.ClienteId) (*internal.Cliente, error) {
	var cliDbConTaDb clienteDbConTipoAtraccionDb
	err := cl.DB.QueryRowx(ObtenerClienteSql, idCliente).StructScan(cliDbConTaDb)

	if err != nil {
		return nil, err
	}

	cli := clienteDbToCliente(cliDbConTaDb)
	return &cli, nil
}

func (cl *ClienteStorage) Actualizar(cliente *internal.Cliente) error {
	cliDb := clienteToClienteDb(*cliente)
	panic("not implemented")
}

func (cl *ClienteStorage) Crear(cliente *internal.Cliente, pwd string) error {
	cliDb := clienteToClienteDb(*cliente)
	cliPwdDb := generateClientePasswordDb(int(cliente.Id), cliente.Username, pwd)

	tx := cl.DB.MustBegin()
	tx.MustExec(CrearClienteSql, cliDb)
	tx.MustExec(CrearPasswordCliente, cliPwdDb)
	err := tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (cl *ClienteStorage) Validar(username, pwd string) error {
	panic("Not implemented")
}

func (cl *ClienteStorage) ActualizarPassword(username, pwdAnterior, pwdNueva, pwdNuevaConfirmacion string) error {
	if pwdNueva != pwdNuevaConfirmacion {
		return errors.New("Las contraseña nueva y su confirmación no coinciden entre si")
	}

	if len(pwdNueva) < 8 && len(pwdNueva) > 16 {
		return errors.New("El largo de la contraseña debe estar entre 8 y 16 caracteres")
	}

	panic("not implemented")
}

// Mappers de tabla cliente
type clienteDb struct {
	cli_id          int
	cli_usrnm       string
	cli_din_disp    decimal.Decimal
	cli_tiempo_disp time.Duration
	cli_tip_atr_id  int16
	cli_fec_alt     sql.NullTime
	cli_fec_act     sql.NullTime
	cli_fec_baj     sql.NullTime
}

type clienteDbConTipoAtraccionDb struct {
	clienteDb
	tipoAtraccionDb
}

func clienteDbToCliente(cliDbConTaDb clienteDbConTipoAtraccionDb) internal.Cliente {
	cli := internal.Cliente{
		Id:               internal.ClienteId(cliDbConTaDb.cli_id),
		Username:         cliDbConTaDb.cli_usrnm,
		Dinero:           cliDbConTaDb.cli_din_disp,
		TiempoDisponible: cliDbConTaDb.cli_tiempo_disp,
		AtraccionPreferida: &internal.TipoAtraccion{
			Id:        internal.TipoAtraccionId(cliDbConTaDb.tip_atr_id),
			Nombre:    cliDbConTaDb.tip_atr_nom,
			FechaAlta: cliDbConTaDb.tip_atr_fec_alt.Time,
			FechaBaja: nil,
		},
		FechaAlta:          cliDbConTaDb.cli_fec_alt.Time,
		FechaActualizacion: cliDbConTaDb.cli_fec_act.Time,
		FechaBaja:          nil,
	}

	if cliDbConTaDb.cli_fec_baj.Valid {
		cli.FechaBaja = &cliDbConTaDb.cli_fec_baj.Time
	}

	if cliDbConTaDb.tip_atr_fec_baj.Valid {
		cli.AtraccionPreferida.FechaBaja = &cliDbConTaDb.cli_fec_baj.Time
	}

	return cli
}

func clienteToClienteDb(cliente internal.Cliente) clienteDb {
	return clienteDb{
		cli_id:          int(cliente.Id),
		cli_usrnm:       cliente.Username,
		cli_din_disp:    cliente.Dinero,
		cli_tiempo_disp: cliente.TiempoDisponible,
		cli_tip_atr_id:  int16(cliente.AtraccionPreferida.Id),
		cli_fec_alt:     sql.NullTime{Time: cliente.FechaAlta},
		cli_fec_act:     sql.NullTime{Time: cliente.FechaActualizacion},
		cli_fec_baj:     sql.NullTime{Time: *cliente.FechaBaja},
	}
}
