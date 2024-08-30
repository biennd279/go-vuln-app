// Package web_mux internal/adapter/web_mux/httpfx.go
package web_mux

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Module("web_mux",
	fx.Provide(
		NewCommandInjectionAdapter,
		NewHttpApplication,
	),
)

func NewHttpApplication(adapter *CommandInjectionAdapter) *mux.Router {
	r := mux.NewRouter()
	adapter.RegisterRoutes(r)
	return r
}
