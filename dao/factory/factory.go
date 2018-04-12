package factory

import (
	"../interfaces"
	"../mysql"
	"log"
)


func FactoryDaoUsuario(e string) interfaces.UsuarioDao {
	var i interfaces.UsuarioDao

	switch e {
	case "mysql":
		i = mysql.MysqlImplDb{}

	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}
