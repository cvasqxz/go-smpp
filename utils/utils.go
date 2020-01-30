package utils

import (
	"encoding/binary"
	"log"
)

func Int2byte(i uint32) []byte {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, i)
	return buffer
}

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
