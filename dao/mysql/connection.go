package mysql

import (
	core "github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//fmt "fmt"
	"../../utilities"
	"log"
//models "../../models"
	//"sync"
	//"fmt"
)





func getConection() *xorm.Engine {

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	engine, err = xorm.NewEngine("mysql", config.User+":"+config.Password+"@tcp("+config.Server+":"+config.Port+")/"+config.Database+"?charset=utf8&parseTime=True")
	engine.SetMapper(core.SameMapper{})
	//engine.Sync2(new(models.Radicacion))



	/*dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}*/
	return engine
}
