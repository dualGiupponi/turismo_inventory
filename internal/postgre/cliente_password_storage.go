package postgre

import "time"

const (
	CrearPasswordCliente string = `INSERT INTO clientes_passwords (clpwd_cli_id, clpwd_usrnm, clpwd_pwd, clpwd_seed, clpwd_fec_act) 
									VALUES(@clpwd_cli_id, @clpwd_usrnm, @clpwd_pwd, @clpwd_seed, @clpwd_fec_act);`
	ObtenerSeedPwdClienteSql = "SELECT clpwd_seed FROM clientes_passwords WHERE clpwd_usrnm = @clpwd_usrnm"
	ValidarPwdClienteSql     = "SELECT count(*) FROM clientes_passwords WHERE clpwd_usrnm = @clpwd_usrnm AND clpwd_pwd = @clpwd_pwd"
)
)

type clientePasswordDb struct {
	clpwd_cli_id  int
	clpwd_usrnm   string
	clpwd_pwd     string
	clpwd_seed    string
	clpwd_fec_act time.Time
}

func generateClientePasswordDb(cliId int, username, password string) clientePasswordDb {
	encryptPwd, seed := encryptPassword(password)

	return clientePasswordDb{
		clpwd_cli_id:  cliId,
		clpwd_usrnm:   username,
		clpwd_pwd:     encryptPwd,
		clpwd_seed:    seed,
		clpwd_fec_act: time.Now(),
	}
}
