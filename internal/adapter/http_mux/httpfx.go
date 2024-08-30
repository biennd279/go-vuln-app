// Package http_mux internal/adapter/http_mux/httpfx.go
package http_mux

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var Module = fx.Module("http_mux",
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
