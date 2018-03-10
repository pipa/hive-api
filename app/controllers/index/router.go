package indexController

import (
	"github.com/kataras/iris/core/router"
	"github.com/pipa/hive-api/app/utils"
)

// EquipRouter will add the router and run the handler for each of the routes
func EquipRouter(app router.Party) {
	party := app.Party("/")

	for _, route := range routes {
		party.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}

// routes all the routes for this controller
var routes = utils.Routes{
	utils.Route{
		Method:      "GET",
		Pattern:     "/ip/{ip}",
		HandlerFunc: Index,
	},
}
