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
