package utils

import (
	"github.com/kataras/iris/context"
)

type (
	// Routes is a slice of Route, used for creating routes in a controller
	Routes []Route

	// Route is the basic structure of a Route
	Route struct {
		Method      string
		Pattern     string
		HandlerFunc context.Handler
	}
)
