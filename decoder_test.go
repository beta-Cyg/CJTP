package cjtp

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	encoded_bytes, err := Encode(CJTPSegment{0, 10, 0, 100, []byte("test")})
	if err != nil {
		panic(err)
	}
	decoded_segment, err := Decode(encoded_bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", decoded_segment)
}

