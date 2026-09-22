package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

func (app *application) serve() error {
	app.infoLog.Println("API, listing on port", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	var cgf config
	cgf.port = 8080
	fmt.Printf("Config: %+v\n", cgf)
	app := &application{
		config:   cgf,
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
