package internal

import (
	"github.com/praveenmahasena/libsys/internal/server"
)

func Start() error {

	ns := server.New(":42069")

	return ns.Run()
}
