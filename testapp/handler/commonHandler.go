package handler

import (
	"AetherGo/internal/context"
	"AetherGo/internal/render"
)

// func IndexHandler(c *context.Context) {
// 	render.RenderJSON(c.Response, map[string]string{"message": "Hello, World!"})
// }

func IndexHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/index.html", nil)
	}
}

func RegisterHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/register.html", nil)
	}
}

func DashboardHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/dashboard.html", nil)
	}
}

func TransactionHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/transactions.html", nil)
	}
}

func SendMoneyHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/send-money.html", nil)
	}
}

func CashInHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/cash-in.html", nil)
	}
}

func CashOutHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/cash-out.html", nil)
	}
}

func SettingsHandler(c *context.Context) {
	if c.Request.Method == "GET" {
		render.RenderTemplate(c.Response, "testapp/templates/settings.html", nil)
	}
}
