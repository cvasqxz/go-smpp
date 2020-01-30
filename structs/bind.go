package structs

import (
	"encoding/hex"
	"fmt"
	"github.com/cvasqxz/go-smpp/variables"
)

type bind struct {
	header           Header
	systemID         string
	password         string
	systemType       string
	interfaceVersion uint8
	addrTON          uint8
	addrNPI          uint8
	addressRange     []byte
}

type bindRESP struct {
	header   Header
	systemID string
}

func CreateBind(command string, sequence uint32, systemID string, password string, systemType string, addrTON string, addrNPI string) bind {
	pdu := bind{}

	pdu.header.commandID = "bind_" + command
	pdu.header.status = "ROK"
	pdu.header.sequence = sequence

	if len(password) <= 9 {
		pdu.password = password
	}
	if len(systemID) <= 16 {
		pdu.systemID = systemID
	}
	if len(systemType) <= 13 {
		pdu.systemType = systemType
	}

	pdu.addrTON = variables.TypeOfNumber[addrTON]
	pdu.addrNPI = variables.NumericPlanIndicator[addrNPI]

	pdu.header.length = uint32(16 + len(password) + len(systemID) + len(systemType) + 3 + len(pdu.addressRange))

	return pdu
}

func (pdu *bind) PackBind() string {
	var buffer string
	buffer += pdu.header.pack()
	buffer += hex.EncodeToString([]byte(pdu.systemID))
	buffer += hex.EncodeToString([]byte(pdu.password))
	buffer += hex.EncodeToString([]byte(pdu.systemType))
	buffer += fmt.Sprintf("%02x", pdu.interfaceVersion)
	buffer += fmt.Sprintf("%02x", pdu.addrTON)
	buffer += fmt.Sprintf("%02x", pdu.addrNPI)
	buffer += hex.EncodeToString(pdu.addressRange)
	return buffer
}

type unbind struct {
	header Header
}

type unbindRESP struct {
	header Header
}

func (pdu *unbind) PackUnbind() string {
	return pdu.header.pack()
}
