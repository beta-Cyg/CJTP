package cjtp

import "fmt"

type CJTPSegment struct {
	version uint8 `default:0`
	length uint32 `default:0`
	r_flag uint8 `default:0`
	r_code uint8 `default:0`
	data []byte `default:nil`
}

func (segment *CJTPSegment) Format() string {
	var r string
	if segment.r_flag == CJTP_REQUEST {
		r = "Request"
	} else {
		r = "Response"
	}
	return fmt.Sprintf("CJTP Version: %v, Segment Length: %v, %s, Code: %v, Data: %s",
		segment.version,
		segment.length,
		r,
		segment.r_code,
		string(segment.data))
}

func (segment *CJTPSegment) Data2String() string {
	return string(segment.data)
}

const (
	CJTP_RESPONSE uint8 = 0
	CJTP_REQUEST uint8 = 1
)

