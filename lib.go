package glog

import (
	"log"
	"os"
)

// Log Wrapper
type Log struct {
	Mode   LogMode
	Logger *log.Logger
}

// LogMode for vos/session
type LogMode uint8

// LogMode
const (
	No LogMode = iota // No Log
	Error
	Warn
	Info
	Debug
)

// New Log
func New() *Log {
	return &Log{
		Mode:   Info,
		Logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

// Log Action
func (l *Log) Log(prompt string, color func(a ...interface{}) string, vs ...interface{}) {
	vss := make([]interface{}, 0, len(vs)+1)
	vss = append(vss, color("["+prompt+"]"))
	vss = append(vss, vs...)
	l.Logger.Println(vss...)
}

// Debug Log
func (l *Log) Debug(vs ...interface{}) {
	if l.Mode == Debug {
		l.Log("D", Green, vs...)
	}
}

// Info Log
func (l *Log) Info(vs ...interface{}) {
	if l.Mode >= Info {
		l.Log("I", Blue, vs...)
	}
}

// Warn Log
func (l *Log) Warn(vs ...interface{}) {
	if l.Mode >= Warn {
		l.Log("W", Yellow, vs...)
	}
}

// Error Log
func (l *Log) Error(e error) {
	if l.Mode >= Error {
		l.Log("E", Red, e)
	}
	panic(e)
}
