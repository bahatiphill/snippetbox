package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//Define an application struct to hold the application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//Check if there is a "addr" flag passed on the command line and use it as port number
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	//creating loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Initialize a new instance of application containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//initialize a http.server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
