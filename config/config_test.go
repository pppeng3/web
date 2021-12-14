package config

import "testing"

func TestGetConfig(t *testing.T) {
	Init()
	t.Log(GetConfig())
}
