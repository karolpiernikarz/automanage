package newAdmin

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	DB = DBConnect()
}

func DBConnect() *gorm.DB {
	db, err := sqlx.Open("mysql", viper.GetString("newadmin.db.username")+":"+viper.GetString("newadmin.db.password")+"@tcp("+viper.GetString("newadmin.db.host")+":"+viper.GetString("newadmin.db.port")+")/"+viper.GetString("newadmin.db.database"))
	if err != nil {
		fmt.Println(err)
	}

	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	return gormdb
}
