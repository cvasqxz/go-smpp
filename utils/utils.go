package utils

import (
	"bufio"
	"encoding/binary"
	"log"
	"net"
)

func Int2byte(i uint32) []byte {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, i)
	return buffer
}

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Byte2int(slice []byte) uint32 {
	data := uint32(0)
	for _, b := range slice {
		data = (data << 8) | uint32(b)
	}
	return data
}

func OpenConn(addr string) (*bufio.ReadWriter, error) {
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), err
}
