package interfaces

import (
	"../../models"
	"database/sql"
)

type UsuarioDao interface {
	CreateUsuario(u models.Usuarios) (int64,error)
	FindUsuarioLogin(id string,pass string)(sql.NullInt64, error)
	GetAllUsuario() ([]models.Usuarios, error)
	GetAllUsuarioExtend() ([]models.UsuariosSedes, error)
}
