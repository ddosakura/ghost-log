package glog

import (
	"errors"
)

type obj struct {
	*Log
}

func Example() {
	defer func() {
		_ = recover()
	}()

	o := &obj{New()}
	//o.Logger = log.New(os.Stderr, "[Example] ", log.LstdFlags)
	o.Logger.SetPrefix("[Example] ")
	o.Mode = Debug
	o.Debug("Hello World")
	o.Info("Hello World")
	o.Warn("Hello World")
	o.Error(errors.New("Hello World"))

	// Output:
	//
}
