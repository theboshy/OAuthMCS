

package models

import (
	"database/sql"
	"time"
)

type Sedes struct {
	IdSedes int `xorm:"int(11) primary key not null 'idsedes'"`
	NombreSede string `xorm:"VARCHAR(60) 'nombre_sede'"`
	PrefijoSede string `xorm:"VARCHAR(6) 'prefijo_sede'"`
	GmtSede string `xorm:"VARCHAR(45) 'gmt_sede'"`
	DireccionSede string `xorm:"VARCHAR(80) 'direccion_sede'"`
    TelefonoSede string `xorm:"VARCHAR(60) 'telefono_sede'"`
    EstadoSede sql.NullInt64 `xorm:"TINYINT(1) 'estado_sede'"`
    FechaCreacion time.Time `xorm:"TIMESTAMP 'fecha_creacion' null"`
    PrefijoRecibida string `xorm:"VARCHAR(10) 'prefijo_recibida'"`
    ConsecRecibida sql.NullInt64 `xorm:"TINYINT(11) 'consec_recibida'"`
    CiudadesIdCiudades sql.NullInt64 `xorm:"TINYINT(11) 'ciudades_idCIUDADES'"`
}

