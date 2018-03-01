package main

import (
	"flag"
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
	fPort       string
)

func main() {

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

	fScript := flag.Bool("logScript", false, FLAG_SCRIPT)
	fTime := flag.Bool("logTime", false, FLAG_TIME)
	flag.StringVar(&fPort, "port", "9999", FLAG_PORT)
	flag.Parse()

	log.Printf(FLAGGED_WITH, "fScript", *fScript)
	log.Printf(FLAGGED_WITH, "fTime", *fTime)
	log.Printf(FLAGGED_WITH, "fPort", fPort)

	initLog(*fScript, *fTime)

	application = App{}
	application.Initialize()
	application.Run()
}

// Initialize the application
func (a *App) Initialize() {
	a.Router = NewRouter(routes)
	a.Version = "V1.00"
	a.Port = ":" + fPort
}

// Starts the application
func (a *App) Run() {
	log.Printf(API_RUNNING_ON, a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
