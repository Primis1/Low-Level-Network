package logging

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
)

type logLevel int

const (
	INFO logLevel = iota
	ERR
)

type application struct {
	level    logLevel
	errorLog *log.Logger
	infoLog  *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO: \t", log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR: \t", log.Ltime)

var path string

func NewLogger(Level logLevel) *application {
	return &application{
		level:    Level,
		errorLog: errorLog,
		infoLog:  infoLog,
	}
}

func (l *application) Info(msg any, args ...any) {
	if l.level <= INFO {
		file, line := getCaller()
		formattedMessage := fmt.Sprintf("" + msg, args...) // Format the message using the provided arguments
		l.infoLog.Printf("[%s : %d] \n\n%s\n\n", file, line, formattedMessage)
	}
}

func (l *application) Error(err ...any) {
	if l.level <= ERR {
		file, line := getCaller()
		l.errorLog.Fatalf("[%s : %d] \n\n%s\n\n", file, line, fmt.Sprint(err...))
	}
}

func getCaller() (string, int) {
	key := os.Getenv("KEY_WORD") // assign your folder name, so we can crop no-reqired part

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		log.Fatal("runtime caller has an error")
	}

	if key == "" {
		fmt.Print("key is empty")
		return file, line // Return without modifying if key is not set
	}

	regExp, _ := regexp.Compile(".*" + regexp.QuoteMeta(key)) // regex for deleting left side 

	file = regExp.ReplaceAllString(file, key)

	return file, line
}
