package mysql

import "../../models"

func (dao MysqlImplDb) FindBySedes (idSedes int) (models.Sedes,error){
	db := getConection()
	defer db.Close()
	radii := models.Sedes{}
	var affect bool
	var error error

	affect, error = db.Where("idsedes = ?", idSedes).Get(&radii)

	if error != nil || affect == false{
     return radii,error
	}
	return radii,nil

}