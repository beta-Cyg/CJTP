package cjtp

import (
	"net"
	"fmt"
	"testing"
)

func TestReadSegment(t *testing.T) {
	c := make(chan bool)
	sent_seg := CJTPSegment{0, 10, 0, 100, []byte("seed")}
	go func() {
		listener, err := net.Listen("tcp", "127.0.0.1:1145")
		if err != nil {
			panic(err)
		}
		defer listener.Close()
		c <- true
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		encoded_bytes, err := Encode(sent_seg)
		conn.Write(encoded_bytes)
	}()
	<- c
	conn, err := net.Dial("tcp", "127.0.0.1:1145")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	seg, err := ReadSegment(conn)
	if err != nil {
		panic(err)
	}
	seg_bytes, err := Encode(seg)
	if err != nil {
		panic(err)
	}
	sent_seg_bytes, err := Encode(sent_seg)
	if err != nil {
		panic(err)
	}
	if string(seg_bytes) != string(sent_seg_bytes){
		panic(fmt.Errorf("mismatch segment"))
	}
	fmt.Printf("%v\n", seg.Format())
}
