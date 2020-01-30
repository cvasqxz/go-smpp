package structs

import (
	"encoding/hex"
	"github.com/cvasqxz/go-smpp/variables"
	"github.com/cvasqxz/go-smpp/utils"
)

type Header struct {
	length    uint32
	commandID string
	status    string
	sequence  uint32
}

func (header *Header) Pack() string {
	var buffer string

	buffer += hex.EncodeToString(utils.Int2byte(header.length))
	buffer += hex.EncodeToString(variables.Commands[header.commandID])
	buffer += hex.EncodeToString(variables.Status[header.status])
	buffer += hex.EncodeToString(utils.Int2byte(header.sequence))
	return buffer
}
