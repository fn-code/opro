package opro

import (
	"testing"
)

func TestCheckHeader(t *testing.T) {
	var tb = []struct {
		val    []byte
		expect error
	}{
		{[]byte{}, errEmptyByte},
		{[]byte{0x81, 0x80 | byte(32), 0x7d}, nil},
		{[]byte{0x81, 0x8, 0x7d}, errMaskInvalid},
		{[]byte{0x80, 0x80 | byte(32), 0x7d}, errModeInvalid},
	}

	for _, ts := range tb {
		err := checkHeader(ts.val)
		if err != ts.expect {
			t.Errorf("error got %v, expect %v", err, ts.expect)
		}
	}

}
