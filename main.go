package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router  *mux.Router
	Port    string
	Version string
}

var (
	application App
	sPort       string
	logger      log.Logger
)

func StartApi(log log.Logger, fScript bool, fTime bool, fPort string) {
	logger = log
	// this comes from http://www.kammerl.de/ascii/AsciiSignature.php
	// the font used id "standard"
	logger.Println(`                                    `)
	logger.Println(` _                                  `)
	logger.Println(`| |    __ _  __ _  ___   ___  _ __  `)
	logger.Println(`| |   / _  |/ _  |/ _ \ / _ \| '_ \ `)
	logger.Println(`| |__| (_| | (_| | (_) | (_) | | | |`)
	logger.Println(`|_____\__,_|\__, |\___/ \___/|_| |_|`)
	logger.Println(`            |___/                   `)

	logger.Println(`    _    ____ ___                   `)
	logger.Println(`   / \  |  _ \_ _|                  `)
	logger.Println(`  / _ \ | |_) | |                   `)
	logger.Println(` / ___ \|  __/| |                   `)
	logger.Println(`/_/   \_\_|  |___|                  `)
	logger.Println(`                                    `)
	logger.Println(`                                    `)

	logger.Printf(FLAGGED_WITH, "fScript", fScript)
	logger.Printf(FLAGGED_WITH, "fTime", fTime)
	logger.Printf(FLAGGED_WITH, "fPort", fPort)

	sPort = fPort
	initLog(fScript, fTime)

	application = App{}
	application.initialize()
	application.run()
}

// Initialize the application
func (a *App) initialize() {
	a.Router = NewRouter(routes)
	a.Version = "V1.00"
	a.Port = ":" + sPort

}

// Starts the application
func (a *App) run() {
	logger.Printf(API_RUNNING_ON, a.Port)
	logger.Fatal(http.ListenAndServe(a.Port, a.Router))
}
