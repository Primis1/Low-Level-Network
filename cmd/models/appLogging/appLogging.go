package applogging

import (
	"log"
	"os"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO: \t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR: \t", log.Ltime|log.Llongfile)

var App = &Application{
	errorLog: errorLog,
	infoLog:  infoLog,
}