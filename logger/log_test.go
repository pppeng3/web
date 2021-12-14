package logger

import "testing"

func TestLogger(t *testing.T) {
	InitLog()
	Info("Testing")
}
