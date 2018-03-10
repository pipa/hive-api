package middlewares

import (
	"github.com/kataras/iris"
	"github.com/pipa/hive-api/middlewares/jwtAuth"
	"github.com/pipa/hive-api/middlewares/recover"
)

// Wrap adds app middlewares to be used by every request
func Wrap(app *iris.Application) {
	app.Use(recover.New())
	app.Use(jwtAuth.Serve)
}
