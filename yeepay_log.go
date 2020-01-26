package yeepayGo

import (
	"log"
)

type YeepayLoggerInterface interface {
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Panic(args ...interface{})
}

type YeepayLogger struct {
}

func NewYeepayLogger() YeepayLoggerInterface {
	return &YeepayLogger{}
}

func (logs *YeepayLogger) Error(args ...interface{}) {
	log.Fatal(args)
}

func (logs *YeepayLogger) Info(args ...interface{}) {
	log.Println("INFO",args)
}

func (logs *YeepayLogger) Debug(args ...interface{}) {
	log.Println("DEBUG",args)
}

func (logs *YeepayLogger) Warn(args ...interface{}) {
	log.Println("WARN",args)
}

func (logs *YeepayLogger) Panic(args ...interface{}) {
	log.Panic(args)
}
