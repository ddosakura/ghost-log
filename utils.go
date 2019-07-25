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
	logFile, err := os.OpenFile(dir+fileName+".log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return nil
	}
	w := io.MultiWriter(os.Stderr, logFile)
	return log.New(w, "["+fileName+"] ", flag)
}
