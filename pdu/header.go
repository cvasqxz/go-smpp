package pdu

import (
	"encoding/hex"

	"github.com/cvasqxz/go-smpp/utils"
	"github.com/cvasqxz/go-smpp/variables"
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

func (header *Header) GetCommandID() string {
	return header.commandID
}
func (header *Header) GetSequence() uint32 {
	return header.sequence
}

func ParseHeader(bin_packet []byte) Header {
	header := Header{}

	header.length = utils.Byte2int(bin_packet[0:4])
	header.commandID = variables.CommandsHex[hex.EncodeToString(bin_packet[4:8])]
	header.status = variables.StatusHex[hex.EncodeToString(bin_packet[8:12])]
	header.sequence = utils.Byte2int(bin_packet[12:16])

	return header
}
