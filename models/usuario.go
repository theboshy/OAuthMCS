package models

import (
	"time"
	"./custom"
)


type Usuarios struct {
	IdUsuarios int `xorm:"int(11) primary key not null 'idUSUARIOS'" `
	NombreIdUsuario string `xorm:"varchar(15)  'nom_id_usuario'"`
	NombreUsuario string `xorm:"varchar(45) 'nombre_usuario'"`
	FechaCreacionUsuario time.Time `xorm:"TIMESTAMP  'fec_crea_usuario'"`
	EstadoUsuario custom.JsonNullInt64 `xorm:"TINYINT(1) not null 'estado_usuario'"`
	Contrasena string `xorm:"TEXT 'contrasena'"`
	Email string `xorm:"TEXT 'email'"`
	Conexion custom.JsonNullInt64 `xorm:"TINYINT(1) 'Conexion'"`
	Telefono string `xorm:"VARCHAR(45) 'telefono'"`
	SedesIdSedes custom.JsonNullInt64  `xorm:"INT(11) not null 'sedes_idsedes'"`
}

type UsuariosSedes struct {
	Usuarios `xorm:"extends"`
	Sedes `xorm:"extends"`
}

func (UsuariosSedes) TableName() string {
	return "usuarios"
}
