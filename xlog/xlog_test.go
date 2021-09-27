package xlog

import "testing"

func TestDebug(t *testing.T) {
	Debug("xxxxx")
}

func TestDebugf(t *testing.T) {
	Debugf("TestDebugf:%v", "args...")
}

func TestInfo(t *testing.T) {
	Info("args...")
}

func TestInfof(t *testing.T) {
	Infof("TestInfof%v", "args...")
}

func TestWarn(t *testing.T) {
	Warn("args...")
}

func TestWarnf(t *testing.T) {
	Warnf("TestWarnf%v", "args...")
}

func TestError(t *testing.T) {
	Error("args...")
}

func TestErrorf(t *testing.T) {
	Errorf("TestErrorf%", "args...")
}

func TestDPanic(t *testing.T) {
	DPanic("args...")
}

func TestDPanicf(t *testing.T) {
	DPanicf("TestDPanicf", "args...")
}

func TestFatal(t *testing.T) {
	//Fatal("args...")
}

func TestFatalf(t *testing.T) {
	//Fatalf("TestFatalf%v", "args...")
}

func TestPanicf(t *testing.T) {
	//Panicf("TestPanicf%v", "args...")
}

func TestPanic(t *testing.T) {
	//Panic("args...")
}