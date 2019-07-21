package initiator

import (
	"github.com/liweiyuan/go-learn/infra/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var INSTANCE *gorm.DB

func init() {

	//db.init
	connectString := config.GetMysqlGreConfig()
	fmt.Printf(connectString)

	connect, err := gorm.Open("mysql", connectString)
	connect.LogMode(true)
	if err != nil {
		fmt.Println(err)
		panic("connect mysql failed")
	}
	INSTANCE = connect
}
