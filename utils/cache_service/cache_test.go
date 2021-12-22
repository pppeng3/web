package cache_service

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	Instance().Set("ropz", "ropz2", 2)
	t.Log(Instance().Get("ropz"))
	time.Sleep(time.Second * 2)
	t.Log(Instance().Get("ropz"))
}
