package config

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ReadEnv(key, defVal string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		//log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", key, defVal)
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}


func ConnectionDB() (*sql.DB, error) {

	dbUser := ReadEnv("dbUser", "root")
	dbPass := ReadEnv("dbPass", "Viontin@12")
	dbHost := ReadEnv("dbHost", "localhost")
	dbPort := ReadEnv("dbPort", "3306")
	dbName := ReadEnv("dbName", "db_hotel")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	}

	return db, nil
}