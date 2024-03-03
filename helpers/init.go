package helpers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var gormDB *gorm.DB

func Init() {
	// init database
	db, err := sqlx.Open("mysql", viper.GetString("db.username")+":"+viper.GetString("db.password")+"@tcp("+viper.GetString("db.host")+":"+viper.GetString("db.port")+")/")

	// set max connection
	db.SetMaxOpenConns(50)

	if err != nil {
		log.Fatal(err)
	}
	// check if connection is working
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// init gorm
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// silent gorm
	gormDB.Logger = gormDB.Logger.LogMode(logger.Error)
}
