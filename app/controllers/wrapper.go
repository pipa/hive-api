package controllers

import (
	"github.com/kataras/iris"
	IndexController "github.com/pipa/hive-api/app/controllers/index"
)

// WithRouter wraps the iris app and calls the controllers in this directory
func WithRouter(app *iris.Application) {
	mainRouter := app.Party("/")

	IndexController.EquipRouter(mainRouter)
}
