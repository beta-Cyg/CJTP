package cjtp

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	seg := CJTPSegment{0, 10, CJTP_REQUEST, 100, []byte("test")}
	fmt.Printf("%s\n", seg.Format())
}

func TestData2String(t *testing.T) {
	seg := CJTPSegment{0, 12, CJTP_REQUEST, 100, []byte("cygnus")}
	if seg.Data2String() != "cygnus" {
		panic(fmt.Errorf("mismatch data string"))
	}
	fmt.Printf("%s\n", seg.Data2String())
}
