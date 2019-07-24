package glog

import (
	"io"
	"log"
	"os"
)

// StopPanic util
func StopPanic() {
	_ = recover()
}

// MultiLogger Builder
func MultiLogger(dir, fileName string, flag int) *log.Logger {
	logFile, err := os.Open(dir + fileName + ".log")
	if err != nil {
		return nil
	}
	w := io.MultiWriter(os.Stderr, logFile)
	return log.New(w, "["+fileName+"] ", flag)
}
