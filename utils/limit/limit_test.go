package limit

import (
	"testing"
)

func TestAddLimitData(t *testing.T) {
	AddLimitData("ropz2", "000")
	AddLimitData("ropz2", "000")
}

func TestHasOverflow(t *testing.T) {
	assertEqual(t, HasOverflow("ropz2", "000"), true)
	AddLimitData("ropz2", "000")
	assertEqual(t, HasOverflow("ropz2", "000"), true)
	AddLimitData("ropz2", "000")
	assertEqual(t, HasOverflow("ropz2", "000"), false)
}

func assertEqual(t *testing.T, a, b interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}

func TestCheckUpperLimit(t *testing.T) {
	assertEqual(t, CheckUpperLimit("ropz2", "0000"), true)
	assertEqual(t, CheckUpperLimit("ropz2", "0000"), false)
	assertEqual(t, CheckUpperLimit("ropz2", "0000"), false)
}
