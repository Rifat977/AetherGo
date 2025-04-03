package app

import (
	"net/http"
	"sync"

	"AetherGo/internal/config"
	"AetherGo/internal/db"
	"AetherGo/internal/log"
	"AetherGo/internal/middleware"
	"AetherGo/internal/router"

	"fmt"
	"os"
	"path/filepath"
)

type App struct {
	Config      *config.Config
	Router      *router.Router
	middlewares []middleware.MiddlewareFunc
	mu          sync.RWMutex
}

func NewApp() *App {

	app := &App{
		Config: config.NewConfig(),
		Router: router.NewRouter(),
	}

	db.ConnectDB()

	return app
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

// CreateNewProject initializes a new project with the required folder structure
func CreateNewProject(projectName string) error {
	// Define the project structure
	dirs := []string{
		projectName,
		filepath.Join(projectName, "cmd"),
		filepath.Join(projectName, "internal"),
		filepath.Join(projectName, "internal/app"),
		filepath.Join(projectName, "internal/db"),
		filepath.Join(projectName, "internal/config"),
		filepath.Join(projectName, "internal/log"),
		filepath.Join(projectName, "internal/router"),
		filepath.Join(projectName, "routes"),
		filepath.Join(projectName, "handler"),
		filepath.Join(projectName, "models"),
		filepath.Join(projectName, "templates"),
	}

	// Create directories
	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	// Create a sample `main.go` file
	mainFilePath := filepath.Join(projectName, "main.go")
	mainFileContent := `package main

import (
	"AetherGo/internal/app"
)

func main() {
	app := app.NewApp()
	app.Run()
}`

	err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to create main.go: %v", err)
	}

	fmt.Printf("Project '%s' created successfully!\n", projectName)
	return nil
}
