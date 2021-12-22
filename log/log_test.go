package log

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	Init("./log", "pp", "pp", "debug")
	Info("Testing")
	time.Sleep(29 * time.Second)
	Debug("DEBUG")
}
