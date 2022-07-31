package util

import "testing"

func TestRandomFixedString(t *testing.T) {
	l := 16
	pw := RandomFixedString(l)
	if len(pw) != l {
		t.Fatal("invalid len", len(pw))
	}
	t.Log(pw)
}
