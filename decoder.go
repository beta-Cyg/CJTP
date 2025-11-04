package cjtp

import (
	"bytes"
	"encoding/binary"
)

func Decode(segment []byte) (CJTPSegment, error) {
	seg := CJTPSegment{0, 0, 0, 0, nil}
	buf := bytes.NewBuffer(segment)
	err := binary.Read(buf, binary.BigEndian, &seg.version)
	if err != nil {
		return CJTPSegment{}, err
	}
	err = binary.Read(buf, binary.BigEndian, &seg.length)
	if err != nil {
		return CJTPSegment{}, err
	}
	var flag_code uint8
	err = binary.Read(buf, binary.BigEndian, &flag_code)
	seg.r_flag = flag_code >> 7
	seg.r_code = (flag_code << 1) >> 1
	seg.data = make([]byte, seg.length - 6/*length of header*/)
	err = binary.Read(buf, binary.BigEndian, &seg.data)
	if err != nil {
		return CJTPSegment{}, err
	}
	return seg, nil
}

