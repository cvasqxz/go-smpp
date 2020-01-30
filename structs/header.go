package structs

import (
	"github.com/cvasqxz/go-smpp/variables"
	"github.com/cvasqxz/go-smpp/utils"
)

type Header struct {
	length    uint32
	commandID string
	status    string
	sequence  uint32
}

func (header *Header) Pack() []byte {
	var buffer []byte

	buffer = append(buffer, utils.Int2byte(header.length)...)
	buffer = append(buffer, variables.Commands[header.commandID]...)
	buffer = append(buffer, variables.Status[header.status]...)
	buffer = append(buffer, utils.Int2byte(header.sequence)...)
	return buffer
}
