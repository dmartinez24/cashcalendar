package app

import (
	"net/http"

	"cashcalendar/app/actions"
	"cashcalendar/app/actions/home"
	"cashcalendar/app/middleware"
	"cashcalendar/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)
	clientsResource := actions.ClientsResource{}
	root.GET("/clients", clientsResource.List).Name("clientList")
	root.ServeFiles("/", http.FS(public.FS()))
}
