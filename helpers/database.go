package helpers

import (
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ImportDatabase(database string, password string) (x string, err error) {
	command := `mysql -h ` + viper.GetString("db.host") + " -u" + viper.GetString("db.username") + " -p" + viper.GetString("db.password") + " " + database + " < web_template.sql"
	cmd := exec.Command("sh", "-c", command)
	err = cmd.Run()
	if err != nil {
		log.WithFields(log.Fields{"database": database, "database_host": viper.GetString("db.host")}).Error("error importing database")
		return "", err
	}
	return
}

func CheckDatabaseUser(username string) (bool, error) {
	var count int64
	err := gormDB.Raw("SELECT COUNT(*) FROM mysql.user WHERE User = ? AND Host = ?", username, "%").Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateDatabase(database string) (err error) {
	// create database
	result := gormDB.Exec("CREATE DATABASE " + database + ";")
	// check if error
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateDatabaseUser(username string, password string) (err error) {
	// create user
	result := gormDB.Exec("CREATE USER '" + username + "'@'%' IDENTIFIED BY '" + password + "';")
	// check if error
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GivePermToUser(username string, database string) (err error) {
	// grant privileges
	result := gormDB.Exec("GRANT ALL PRIVILEGES ON " + database + ".* TO '" + username + "'@'%';")
	// check if error
	if result.Error != nil {
		return result.Error
	}
	return nil
}
