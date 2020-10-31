package server

import (
	"log"
	"os"

	"github.com/bongnv/gwf"
)

// Serve to start a server.
func Serve() error {
	app := gwf.Default()

	s := &Server{
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	v1 := app.Group("/v1")
	v1.POST("/transactions", s.createTransaction)
	v1.GET("/transactions", s.listTransactions)

	return app.Run()
}

// Server ...
type Server struct {
	logger gwf.Logger
}
