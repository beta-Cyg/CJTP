package cjtp

import (
	"bytes"
	"encoding/binary"
)

func Encode(segment CJTPSegment) ([]byte, error) {
	head_buffer := bytes.NewBuffer([]byte{})
	err := binary.Write(head_buffer, binary.BigEndian, segment.version)
	if err != nil {
		return nil, err
	}
	err = binary.Write(head_buffer, binary.BigEndian, segment.length)
	if err != nil {
		return nil, err
	}
	err = binary.Write(head_buffer, binary.BigEndian, segment.r_flag << 7 + segment.r_code)
	if err != nil {
		return nil, err
	}
	err = binary.Write(head_buffer, binary.BigEndian, segment.data)
	if err != nil {
		return nil, err
	}
	return head_buffer.Bytes(), err
}
