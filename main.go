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
)

func StartApi(log log.Logger, fScript bool, fTime bool, fPort string) {

	// this comes from http://www.kammerl.de/ascii/AsciiSignature.php
	// the font used id "standard"
	log.Println(` _                                  `)
	log.Println(`| |    __ _  __ _  ___   ___  _ __  `)
	log.Println(`| |   / _  |/ _  |/ _ \ / _ \| '_ \ `)
	log.Println(`| |__| (_| | (_| | (_) | (_) | | | |`)
	log.Println(`|_____\__,_|\__, |\___/ \___/|_| |_|`)
	log.Println(`            |___/                   `)

	log.Println(`    _    ____ ___                   `)
	log.Println(`   / \  |  _ \_ _|                  `)
	log.Println(`  / _ \ | |_) | |                   `)
	log.Println(` / ___ \|  __/| |                   `)
	log.Println(`/_/   \_\_|  |___|                  `)

	log.Printf(FLAGGED_WITH, "fScript", fScript)
	log.Printf(FLAGGED_WITH, "fTime", fTime)
	log.Printf(FLAGGED_WITH, "fPort", fPort)

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
	log.Printf(API_RUNNING_ON, a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
