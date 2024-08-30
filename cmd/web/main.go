package main

import (
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/adapter/config"
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/adapter/repository"
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/adapter/web"
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/adapter/web_echo"
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/adapter/web_mux"
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/core/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),

		config.Module,
		repository.Module,
		service.Module,
		web.Module,
		web_mux.Module,
		web_echo.Module,

		fx.Provide(
			zap.NewExample,
		),

		fx.Invoke(func(server *http.Server) {}),
	).Run()
}
