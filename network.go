package cjtp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"net"
	"fmt"
)

func ReadSegment(conn net.Conn) (CJTPSegment, error) {
	reader := bufio.NewReader(conn)
	header_buf, err := reader.Peek(6) // header
	if err != nil {
		return CJTPSegment{}, err
	}
	length_buf := bytes.NewBuffer(header_buf[1:5])
	var length uint32
	err = binary.Read(length_buf, binary.BigEndian, &length)
	if err != nil {
		return CJTPSegment{}, err
	}
	if uint32(reader.Buffered()) < length {
		return CJTPSegment{}, fmt.Errorf("Segment buffered is too short: %v bytes", length)
	}
	segment_buf := make([]byte, length)
	readed_length, err := reader.Read(segment_buf)
	if err != nil {
		return CJTPSegment{}, err
	}
	if uint32(readed_length) != length {
		return CJTPSegment{}, fmt.Errorf("Segment readed is too short: %v bytes", readed_length)
	}
	return Decode(segment_buf)
}

