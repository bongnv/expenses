package main

import (
	"os"

	"github.com/bongnv/expenses/backend/server"
)

func main() {
	if err := server.Serve(); err != nil {
		os.Exit(1)
	}
}
