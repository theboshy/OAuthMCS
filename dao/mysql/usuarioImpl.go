package mysql

import (
	"../../models"
"../../utilities"
	"fmt"
	"database/sql"
)

func (dao MysqlImplDb) CreateUsuario(model models.Usuarios) (int64, error) {
	db := getConection()
	defer db.Close()
	radci := model
	affec,error :=db.Insert(&radci)
	if error != nil {
		return affec,error
	}
	return affec,nil
}

func (dao MysqlImplDb) FindUsuarioLogin(id string,pass string) (sql.NullInt64, error) {
	db := getConection()
	defer db.Close()
	radii := models.Usuarios{}
	var affect bool

	var errr  error
		affect, errr = db.Where("nom_id_usuario = ?", id).Get(&radii)
		if affect == true && errr == nil && radii.IdUsuarios !=0{
			validation,errr := utilities.ValidateCrypthUni(pass, radii.Contrasena)
			if validation == true {
				return sql.NullInt64{int64(radii.IdUsuarios),true}, errr
			}
		}
		return sql.NullInt64{0,false},errr
}

func (dao MysqlImplDb) GetAllUsuario() ([]models.Usuarios, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := getConection()
	defer db.Close()
	var radii []models.Usuarios
	error :=db.Find(&radii)
	if error != nil {
		return radii,error
	}
	return radii,nil
}

func (dao MysqlImplDb) GetAllUsuarioExtend() ([]models.UsuariosSedes, error) {
	//query := "SELECT numero_radicacion FROM radicacion"
	//radicEng := make([]models.Radicacion, 0)
	db := getConection()
	defer db.Close()

	var users []models.UsuariosSedes
	err := engine.Join("INNER", "sedes", "sedes.idsedes = usuarios.sedes_idsedes").Find(&users)
	fmt.Println(err)
	return users,err
}



