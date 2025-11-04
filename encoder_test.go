package cjtp

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T){
	encoded_bytes, err := Encode(CJTPSegment{0, 6, CJTP_REQUEST, 100, []byte("")})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", encoded_bytes);
}
