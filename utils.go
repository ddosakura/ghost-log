package glog

// StopPanic util
func StopPanic() {
	_ = recover()
}
