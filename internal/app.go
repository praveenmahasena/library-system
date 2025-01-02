package internal

import (
	"errors"

	db "github.com/praveenmahasena/libsys/internal/database"
	"github.com/praveenmahasena/libsys/internal/server"
	"github.com/spf13/viper"
)

func Start() error {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return errors.New("Error during reading config file")
	}

	dbErr := db.Connect()

	if dbErr != nil {
		return dbErr
	}

	ns := server.New(":42069")

	return ns.Run()
}
