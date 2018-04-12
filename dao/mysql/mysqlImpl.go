package mysql

import (
	"github.com/go-xorm/xorm"
	/*"sync"
	"fmt"*/
)

type MysqlImplDb struct {
}
var engine *xorm.Engine
/*


var once sync.Once

func GetInstance() *xorm.Engine {
	once.Do(func() {
		fmt.Println("Generada nueva conexion")
		engine = getConection()
	})
	return engine
}*/