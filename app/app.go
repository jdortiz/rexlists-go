// Package app provides the front-end application code.
package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"rexlists/config"
)

type App struct {
	serverAddr string
	serverPort uint16
}

func New() App {
	var port int
	var err error
	if port, err = strconv.Atoi(config.GetFromEnvOrDefault("FRONTEND_PORT", "4000")); err != nil {
		port = 4000
	}
	return App{
		serverAddr: config.GetFromEnvOrDefault("FRONTEND_ADDR", "localhost"),
		serverPort: uint16(port),
	}
}

func (app *App) Start() error {
	log.Printf("Launching RexLists front-end server: http://%s:%d\n", app.serverAddr, app.serverPort)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", app.serverAddr, app.serverPort), nil)
}
