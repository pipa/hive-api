package cockroach

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
)

// User is our model, which corresponds to the "accounts" database table.
type User struct {
	ID   int `gorm:"primary_key"`
	name string
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

	// attempt to connect to the mysql database
	db, err := gorm.Open("postgres", dsn)

	// if there was an issue opening the connection, send a 500 error
	if err != nil {
		ctx.Panic()
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
