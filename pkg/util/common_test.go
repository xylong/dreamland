package util

import "testing"

func TestNumToBHex(t *testing.T) {
	t.Logf("%s\n", NumToBHex(1261, 16))
	t.Logf("%s\n", NumToBHex(1261, 36))
}

func TestBHex2Num(t *testing.T) {
	t.Logf("%d\n", BHex2Num("4ED", 16))
	t.Logf("%d\n", BHex2Num("Z1", 36))
}
