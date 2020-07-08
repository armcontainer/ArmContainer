package keystone

import (
	"encoding/hex"
	"testing"
)

func TestVersion(t *testing.T) {
	major, minor := Version()
	t.Log(major, minor)
}

func TestRun(t *testing.T) {
	ks, err := New(ARCH_ARM, MODE_ARM)
	if err != nil {
		t.Fatal(err)
	}
	defer ks.Close()

	// if err := ks.Option(OPT_SYNTAX, OPT_SYM_RESOLVER); err != nil {
	// 	t.Fatal(err)
	// }

	bincode, _, ok := ks.Assemble("str w11, [x13], #0", 0)
	if !ok {
		t.Fatal("assemble error")
	}

	t.Log(hex.EncodeToString(bincode))
}
