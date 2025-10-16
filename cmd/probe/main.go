package main

import (
	"log"
	"os"

	"github.com/belovai/gixer-probe/config"
)

const VERSION = "0.0.1"

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	debugLog *log.Logger
	config   config.Config
}

func main() {
	app := application{}
	err := app.init()
	if err != nil {
		panic(err)
	}

	app.infoLog.Println("gixer-probe v" + VERSION)
	app.debugLog.Printf("app.config: %+v\n", app.config.App)
	app.debugLog.Printf("app.config: %+v\n", app.config.Api)
	app.debugLog.Printf("app.config: %+v\n", app.config.Rmq)

}

func (app *application) init() (err error) {
	app.errorLog = log.New(os.Stderr, "ERROR: ", log.Lshortfile)
	app.infoLog = log.New(os.Stdout, "INFO: ", log.Lshortfile)
	app.debugLog = log.New(os.Stdout, "DEBUG: ", log.Lshortfile)

	app.config, err = config.NewConfig()
	if err != nil {
		return
	}

	return nil
}
