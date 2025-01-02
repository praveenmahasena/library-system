package internal

import (
	"github.com/praveenmahasena/libsys/internal/database"
	"github.com/praveenmahasena/libsys/internal/server"
)

func Start() error {
	dbCon, dbErr := database.Connect()

	if dbErr != nil {
		return dbErr
	}

	ns := server.New(":42069")

	return ns.Run()
}
