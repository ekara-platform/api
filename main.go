package api

import (
	"log"
	"net/http"

	"github.com/ekara-platform/api/secret"
	"github.com/ekara-platform/api/storage"
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
	usedStorage storage.Storage
	usedSecret  secret.Secret
)

func StartApi(log log.Logger, fScript bool, fTime bool, fPort string) {
	logger = log
	// this comes from http://www.kammerl.de/ascii/AsciiSignature.php
	// the font used id "standard"
	log.Println(" _____ _                   ")
	log.Println("| ____| | ____ _ _ __ __ _ ")
	log.Println("|  _| | |/ / _` | '__/ _` |")
	log.Println("| |___|   < (_| | | | (_| |")
	log.Println(`|_____|_|\_\__,_|_|  \__,_|`)

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
	usedStorage = storage.GetStorage()
	usedSecret = secret.GetSecret()
	logger.Printf(API_RUNNING_ON, a.Port)
	logger.Fatal(http.ListenAndServe(a.Port, a.Router))
}
