package app

import (
	"net/http"
	"sync"

	"AetherGo/internal/config"
	"AetherGo/internal/log"
	"AetherGo/internal/middleware"
	"AetherGo/internal/router"
)

type App struct {
	Config      *config.Config
	Router      *router.Router
	middlewares []middleware.MiddlewareFunc
	mu          sync.RWMutex
}

func NewApp() *App {
	return &App{
		Config: config.NewConfig(),
		Router: router.NewRouter(),
	}
}

func (a *App) Use(mw middleware.MiddlewareFunc) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.middlewares = append(a.middlewares, mw)
}

func (a *App) Run() error {
	if a.Config.GetEnv() == "development" {
		log.Printf("Starting development server on http://localhost%s", a.Config.GetPort())
		return http.ListenAndServe(a.Config.GetPort(), a.Router)
	} else {
		log.Printf("Starting production server on http://localhost%s", a.Config.GetPort())
		return http.ListenAndServe(a.Config.GetPort(), a.Router)
	}
}
