package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Load() *sql.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbName := viper.GetString(`database.db_name`)
	dbUser := viper.GetString(`database.db_user`)
	dbPassword := viper.GetString(`database.db_password`)
	dbDriver := viper.GetString(`database.driver`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return dbConn

}
