package main

import (
	"log"

	"github.com/farwydi/demo-cleanarh/provider"
	"go.uber.org/fx"
)

// Консольное приложение
func main() {
	var logger *log.Logger
	app := fx.New(fx.Options(
		provider.DI,

		fx.Populate(&logger),
	))

	if err := app.Err(); err != nil {
		logger.Fatalln(err)
	}

	app.Run()
}
