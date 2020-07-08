package capstone

import (
	"testing"
)

func TestEngine_Dis(t *testing.T) {
	cs, err := New(ARCH_ARM64, MODE_LITTLE_ENDIAN)
	if err != nil {
		t.Fatal(err)
	}
	defer cs.Close()

	codes, err := cs.Dis([]byte("\xab\x05\x00\xb8\xaf\x05\x40\x38"), 0, 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, code := range codes {
		t.Log(code.Mnemonic(), code.OpStr())
	}
}
