package structs

import (
	"github.com/cvasqxz/go-smpp/variables"
)

type bind struct {
	header           Header
	systemID         string
	password         string
	systemType       string
	interfaceVersion []byte
	addrTON          []byte
	addrNPI          []byte
	addressRange     []byte
}

type bindRESP struct {
	header   Header
	systemID string
}

func CreateBind(command string, sequence uint32, systemID string, password string, systemType string, version string, addrTON string, addrNPI string) bind {
	pdu := bind{}

	pdu.header.commandID = "bind_" + command
	pdu.header.status = "ROK"
	pdu.header.sequence = sequence

	if len(password) <= 8 {
		pdu.password = password + "\x00"
	}
	if len(systemID) <= 15 {
		pdu.systemID = systemID + "\x00"
	}
	if len(systemType) <= 12 {
		pdu.systemType = systemType + "\x00"
	}
	pdu.interfaceVersion = variables.InterfaceVersion[version]
	pdu.addrTON = variables.TypeOfNumber[addrTON]
	pdu.addrNPI = variables.NumericPlanIndicator[addrNPI]
	pdu.addressRange = []byte{255}

	pdu.header.length = uint32(16 + len(password) + len(systemID) + len(systemType) + 4)

	return pdu
}

func (pdu *bind) PackBind() []byte {
	var buffer []byte
	buffer = append(buffer, pdu.header.Pack()...)
	buffer = append(buffer, []byte(pdu.systemID)...)
	buffer = append(buffer, []byte(pdu.password)...)
	buffer = append(buffer, []byte(pdu.systemType)...)
	buffer = append(buffer, pdu.interfaceVersion...)
	buffer = append(buffer, pdu.addrTON...)
	buffer = append(buffer, pdu.addrNPI...)
	buffer = append(buffer, pdu.addressRange...)
	return buffer
}

type unbind struct {
	header Header
}

type unbindRESP struct {
	header Header
}

func (pdu *unbind) PackUnbind() []byte {
	return pdu.header.Pack()
}
