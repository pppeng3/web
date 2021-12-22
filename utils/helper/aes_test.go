package helper

import "testing"

func TestAES(t *testing.T) {
	s, _ := AesEncrypt("ppp", "123456789012345612345678")
	t.Log(s)
	// t.Log(AesDecrypt(s, "1234567890123456"))
}
