package cockroach

import (
	"fmt"
	"errors"

	gnorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var (
	db *gnorm.DB
)

type cockroachMiddleware struct {
	config Config
}

// User is our model, which corresponds to the "accounts" database table.
type User struct {
	ID   int `gorm:"primary_key"`
	name string
}

func New(cfg ...Config) context.Handler {
	c := DefaultConfig()
	crMdlw := &cockroachMiddleware{config: c}
	var err error

	// attempt to connect to the mysql database
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbPort, dbHost)
	db, err : gorm.Open("postgres", dsn)

	return crMdlw.ServeHTTP
}

// Serve serves the middleware
func (crMdlw *cockroachMiddleware) ServeHTTP(ctx context.Context) {

}

// Serve creates the jwt-middleware and serves it
func Serve(ctx iris.Context) {
	// get configuration information from the environment
	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbPort := "26257"
	dbName := "hive"

	// build our data source url
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbPort, dbHost)
	ctx.Application().Logger().Debugf("[mdlw-cockroach] dsn ", dsn)

	

	// if there was an issue opening the connection, send a 500 error
	if err != nil {
		ctx.Application().Logger().Debugf("[mdlw-cockroach] Error while startup: ", err)
		ctx.StopExecution()
		return
	}

	// after the middleware finishes, be sure to close our db connection
	defer db.Close()

	// turn off gorm logging
	db.LogMode(false)

	// inject our runtime into the user context for this request
	ctx.Values().Set("db", db)

	// If everything ok then call next.
	ctx.Next()
}
