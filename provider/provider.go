package provider

import (
	"log"
	"os"

	"go.uber.org/fx"
)

var DI = fx.Options(
	// Создание лога
	fx.Provide(func() *log.Logger { return log.New(os.Stdout, "core: ", log.LstdFlags) }),
)
