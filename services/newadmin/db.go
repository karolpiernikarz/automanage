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
	db, err := sqlx.Open("mysql", viper.GetString("db.username")+":"+viper.GetString("db.password")+"@tcp("+viper.GetString("db.host")+":"+viper.GetString("db.port")+")/"+viper.GetString("db.database"))
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
