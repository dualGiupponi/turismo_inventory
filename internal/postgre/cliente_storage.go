package postgre

const (
	ObtenerClienteSql        = "SELECT * FROM clientes WHERE cli_id = @cli_id"
	CrearClienteSql          = ""
	ObtenerSeedPwdClienteSql = ""
	ValidarPwdClienteSql     = ""
)

type ClienteStorage struct {
	*DbConnection
}

func NewClienteStorage(dbConn *DbConnection) *ClienteStorage {
	return &ClienteStorage{DbConnection: dbConn}
}

func (cl *ClienteStorage) Obtener(idCliente ClienteId) (*Cliente, error) {

}

func (cl *ClienteStorage) Actualizar(cliente *Cliente) error {

}

func (cl *ClienteStorage) Crear(cliente *Cliente, pwd string) error {

}

func (cl *ClienteStorage) Validar(username, pwd string) error {

}
