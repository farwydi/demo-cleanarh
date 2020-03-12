package api

import (
	"github.com/farwydi/demo-cleanarh/usecase"
	"go.uber.org/fx"
)

type App struct {
	fx.In

	AuthUseCase usecase.AuthUseCase
}
