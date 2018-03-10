package main

import (
	"github.com/kataras/iris/core/host"
	"github.com/pipa/hive-api/bootstrap"
)

// newApp is used to instantiate the new app
// useful for testing purposes
func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New() // Creating new bootstrap for app

	return app
}

// main is what all executables use to run the app
func main() {
	app := newApp()

	app.ConfigureHost(configureHost)
	app.Listen()
}

// configureHost gives us access to the host created by  `app.Run`,
// they're being executed when application is ready to being served to the public.
func configureHost(host *host.Supervisor) {
	// Register a shutdown "event" callback
	// host.RegisterOnShutdown(func() {})

	// Register a "OnError" callback
	// host.RegisterOnError

	// Register a "OnServe" callback
	// host.RegisterOnServe()
}
