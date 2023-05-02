package logger

import (
	"log"
	"os"
)

type Logging interface {
	Info(msg interface{})
	Infof(msg string, params ...interface{})
	Debug(msg interface{})
	Debugf(msg string, params ...interface{})
	Error(msg interface{})
	Errorf(msg string, params ...interface{})
	Warn(msg interface{})
	Warnf(msg string, params ...interface{})
}

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
	warnLogger  *log.Logger
)

func (l *Logger) Init(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print(err)
	}
	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime)
	warnLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime)

}

func (l *Logger) Info(msg interface{}) {
	infoLogger.Print(msg)

}

func (l *Logger) Infof(msg string, params ...interface{}) {
	infoLogger.Printf(msg, params...)
}

func (l *Logger) Debug(msg interface{}) {
	debugLogger.Print(msg)
}

func (l *Logger) Debugf(msg string, params ...interface{}) {
	debugLogger.Printf(msg, params...)
}

func (l *Logger) Error(msg interface{}) {
	errorLogger.Print(msg)
}

func (l *Logger) Errorf(msg string, params ...interface{}) {
	errorLogger.Printf(msg, params...)
}

func (l *Logger) Warn(msg interface{}) {
	warnLogger.Print(msg)
}

func (l *Logger) Warnf(msg string, params ...interface{}) {
	warnLogger.Printf(msg, params...)
}
