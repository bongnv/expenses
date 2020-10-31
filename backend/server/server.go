package server

import (
	"log"
	"os"

	"github.com/bongnv/expenses/backend/storage"
	"github.com/bongnv/gwf"
	"gorm.io/gorm"
)

// Serve to start a server.
func Serve() error {
	app := gwf.Default()

	db, err := storage.Init()
	if err != nil {
		return err
	}

	s := &Server{
		logger: log.New(os.Stderr, "", log.LstdFlags),
		db:     db,
	}

	v1 := app.Group("/v1")
	v1.POST("/transactions", s.createTransaction)
	v1.GET("/transactions", s.listTransactions)

	return app.Run()
}

// Server ...
type Server struct {
	logger gwf.Logger
	db     *gorm.DB
}
