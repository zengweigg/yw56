package yw56

import (
	"log"
	"os"
)

type Logger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

func createLogger() *logger {
	l := &logger{l: log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)}
	return l
}

var _ Logger = (*logger)(nil)

type logger struct {
	l *log.Logger
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.output("ERROR SfExpress "+format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	l.output("WARN SfExpress "+format, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.output("DEBUG SfExpress "+format, v...)
}

func (l *logger) output(format string, v ...interface{}) {
	if len(v) == 0 {
		l.l.Print(format)
		return
	}
	l.l.Printf(format, v...)
}
