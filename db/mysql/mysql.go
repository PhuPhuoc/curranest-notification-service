package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"githuh.com/PhuPhuoc/curanest-notification-service/config"
)

func ConnectDB() *sqlx.DB {
	dbHost := config.AppConfig.DBHost
	dbPort := config.AppConfig.DBPort
	dbUser := config.AppConfig.DBUser
	dbPassword := config.AppConfig.DBPassword
	dbName := config.AppConfig.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=True&parseTime=True&loc=UTC", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println(dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("error when connect db: ", err)
	}
	return db
}
