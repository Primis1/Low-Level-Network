package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type logLevel int

const (
	INFO logLevel = iota
	ERR
)

type application struct {
	level logLevel
	errorLog *log.Logger
	infoLog  *log.Logger
}

type Args struct {
	fargument string

}

var infoLog = log.New(os.Stdout, "INFO: \t", log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR: \t", log.Ltime)

var path string 

func init() {
	path, _ = os.Getwd()

	fmt.Sprint("\n\n", path)
}

func NewLogger(Level logLevel) *application {
	return &application{
		level: Level,
		errorLog: errorLog,
		infoLog:  infoLog,	
	}
}

func (l *application) Info(msg ...any) {
	if l.level <= INFO {
		file, line := getCaller()
		l.infoLog.Printf("[%s : %d] \n\n%s\n\n",file, line, fmt.Sprint(msg...))
	}
}

func (l *application) Error(err ...any) {
	if l.level <= ERR {
		file, line := getCaller()
		l.errorLog.Fatalf("[%s : %d] \n\n%s\n\n",file, line, fmt.Sprint(err...))
	}
}

func getCaller() (string, int) {
	_, file, line, ok :=  runtime.Caller(2)

	if !ok {
		file = "idk"
		line = 0
	}

	file = strings.TrimPrefix(file, path)
	return file, line
}