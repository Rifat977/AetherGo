package routes

import (
	"AetherGo/internal/app"

	"AetherGo/testapp/handler"
)

func RegisterRoutes(application *app.App) {
	application.Router.Add("GET", "/", handler.IndexHandler)
	application.Router.Add("GET", "/register", handler.RegisterHandler)
	application.Router.Add("GET", "/dashboard", handler.DashboardHandler)
	application.Router.Add("GET", "/transactions", handler.TransactionHandler)
	application.Router.Add("GET", "/send-money", handler.SendMoneyHandler)
	application.Router.Add("GET", "/cash-in", handler.CashInHandler)
	application.Router.Add("GET", "/cash-out", handler.CashOutHandler)
	application.Router.Add("GET", "/settings", handler.SettingsHandler)
}
