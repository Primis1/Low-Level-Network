package applogging

import (
	"log"
	"os"
)

type LogLevel int

type Application struct {
	level LogLevel
	errorLog *log.Logger
	infoLog  *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO: \t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR: \t", log.Ltime|log.Llongfile)


const Level (
	DEBUG LogLevel iota
	INFO
	ERR
)

func NewLogger(Level LogLevel) *Logger {
	return &Application{
		level: Level,
		errorLog: errorLog,
		infoLog:  infoLog,
	
	}
}
