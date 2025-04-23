package routes

import (
	"AetherGo/internal/app"

	"AetherGo/testapp/handler"
)

func RegisterRoutes(application *app.App) {
	application.Router.Add("GET", "/", handler.IndexHandler)
}
