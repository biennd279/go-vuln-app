package web_echo

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Module("web-echo",
	fx.Provide(
		NewCommandInjectionAdapter,
		NewEchoApplication,
	),
)

func NewEchoApplication(adapter *CommandInjectionAdapter) *echo.Echo {
	e := echo.New()
	adapter.RegisterRoutes(e)
	return e
}
