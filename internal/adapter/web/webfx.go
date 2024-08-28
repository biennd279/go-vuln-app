package web

import (
	"context"
	"gitea.homelab.d3s34.me/d3s34/vuln-app/internal/adapter/config"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

const (
	AdapterTag = `group:"adapters"`
)

type GinAdapter interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

var Module = fx.Module("web-adapter",
	fx.Provide(
		AsGinAdapter(NewHealthAdapter),
		fx.Annotate(NewRoutes,
			fx.ParamTags(AdapterTag),
		),
		NewWebApplication,
	),
)

func AsGinAdapter(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(GinAdapter)),
		fx.ResultTags(AdapterTag),
	)
}

func NewRoutes(adapters []GinAdapter, logger *zap.Logger) *gin.Engine {
	r := gin.New()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	r.Use(ginzap.RecoveryWithZap(logger, true))

	v1 := r.Group("/api/v1/")
	for _, adapter := range adapters {
		adapter.RegisterRoutes(v1)
	}

	return r
}

func NewWebApplication(config *config.Config, engine *gin.Engine, lc fx.Lifecycle) *http.Server {

	if !config.IsDev {
		gin.SetMode(gin.ReleaseMode)
	}

	addr := net.JoinHostPort(config.Server.Host, config.Server.Port)

	srv := &http.Server{Addr: addr, Handler: engine} // define a web server

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr) // the web server starts listening on 8080
			if err != nil {

			}
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return srv
}
