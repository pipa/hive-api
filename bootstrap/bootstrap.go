package bootstrap

import (
	stdContext "context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/imdario/mergo"
	"github.com/kataras/iris"
	"github.com/pipa/hive-api/controllers"
	"github.com/pipa/hive-api/middlewares"
	"github.com/pipa/hive-apimiddlewares/onErrorCode"
)

var env = flag.String("env", "dev", "Environment in which the app will run")

// Configurator interface which accepts the framework instance
//
// Used to modify framework instance
type Configurator func(*Bootstrapper)

// Bootstrapper bootstrap the default Iris Application
// with custom methods and properties
type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
func New(cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	// Adding the conf file(merged?) to the iris app instance
	b.Application.Configure(iris.WithConfiguration(b.getConfig()))
	b.AppName = b.Application.ConfigurationReadOnly().GetOther()["appName"].(string)

	b.bootstrap()

	return b
}

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// getConfig parses the `env` flag(if passed) and gets default
// configurations and merges an env conf file.
//
// Also assigns appName and host defaults if not present in the conf
func (b *Bootstrapper) getConfig() iris.Configuration {
	// Flag parsing to have `env` available
	flag.Parse()

	// Default configuration file loaded first
	var confFile iris.Configuration
	defaultFilePath := "./configs/iris.yml"
	if _, err := os.Stat(defaultFilePath); err == nil {
		confFile = iris.YAML(defaultFilePath)
	} else {
		confFile = iris.DefaultConfiguration()
	}

	// If not the default environment
	if *env != "dev" {
		filePath := "./configs/iris." + *env + ".yml"

		if _, err := os.Stat(filePath); err == nil {
			// If environment is not set correctly, app will not start
			envConfFile := iris.YAML(filePath)

			// Merging the default confFile with env file, overriding values
			mergo.Merge(&confFile, envConfFile, mergo.WithOverride)
		} else {
			b.Logger().Debugf("Error with conf file: ", err)
		}
	}

	// Type asserting app name and applying default if not provided
	_, ok := confFile.GetOther()["appName"]
	if !ok {
		confFile.GetOther()["appName"] = "fluffy"
	}

	// Type asserting host address and applying default if not provided
	_, ok = confFile.GetOther()["host"]
	if !ok {
		confFile.GetOther()["host"] = ":8888"
	}

	return confFile
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) bootstrap() *Bootstrapper {
	// All error handlers
	b.setupErrorHandlers()

	// Adding middlewares
	middlewares.Wrap(b.Application)

	// Controllers bootstrap
	controllers.WithRouter(b.Application)

	// Gracefully handle shutdown
	go b.gracefulShutdown(b.Application)

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(cfgs ...iris.Configurator) {
	// Type asserting host address and applying default if not provided
	addr := b.Application.ConfigurationReadOnly().GetOther()["host"].(string)

	b.Run(iris.Addr(addr), cfgs...)
}

// setupErrorHandlers prepares the http error handlers (>=400).
func (b *Bootstrapper) setupErrorHandlers() {
	b.OnAnyErrorCode(onErrorCode.OnAnyErrorCode)
}

// gracefulShutdown handles the shutdown of our app via signals
func (b *Bootstrapper) gracefulShutdown(app *iris.Application) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		// kill -SIGINT XXXX or Ctrl+c
		os.Interrupt,
		syscall.SIGINT, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)
	select {
	case <-ch:
		app.Logger().Info("graceful shutdown...")

		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx)
	}
}
