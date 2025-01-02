package database

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// I really do not like the approach of making the connection as global var
// gotta refactor again

var Con *sqlx.DB

const (
	hostParam     = "HOST"
	portParam     = "PORT"
	userParam     = "USER"
	passwordParam = "PASSWORD"
	dbNameParam   = "DBNAME"
	sslModeParam  = "SSLMODE"
)

func getParams() string {
	host := viper.GetString(hostParam)
	port := viper.GetString(portParam)
	user := viper.GetString(userParam)
	password := viper.GetString(passwordParam)
	dbName := viper.GetString(dbNameParam)
	sslmode := viper.GetString(sslModeParam)
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", host, port, user, password, dbName, sslmode)
}

func Connect() error {
	var err error
	uri := getParams()
	Con, err = sqlx.Connect("postgres", uri)
	if err != nil {
		return errors.New("could not connect to database")
	}
	if err := Con.Ping(); err != nil {
		return errors.New("issue during connection testing")
	}
	return nil
}
