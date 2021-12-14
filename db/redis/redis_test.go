package redis

import "testing"

func TestRedisClient(t *testing.T) {
	rc := Instance(0)
	t.Log(rc)
}

func TestSwitchDB(t *testing.T) {
	rc := Instance(15)
	t.Log(rc)
}
