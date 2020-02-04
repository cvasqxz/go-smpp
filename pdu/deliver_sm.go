package pdu

import "github.com/cvasqxz/go-smpp/variables"

type deliverSM struct {
	header               Header
	serviceType          string
	sourceAddrTON        []byte
	sourceAddrNPI        []byte
	sourceAddr           string
	destAddrTON          []byte
	destAddrNPI          []byte
	destinationAddr      string
	esmClass             []byte
	protocolID           []byte
	priorityFlag         []byte
	scheduleDeliveryTime []byte
	validityPeriod       []byte
	registeredDelivery   []byte
	replaceIfPresentFlag []byte
	dataCoding           []byte
	smDefaultMsgID       []byte
	smLength             []byte
	shortMessage         string
}

type deliverSMRESP struct {
	header    Header
	messageID []byte
}

func CreateDeliverSM(seq uint32, serviceType string,
	sourceAddrTON string, sourceAddrNPI string, sourceAddr string,
	destAddrNPI string, destAddrTON string, destinationAddr string,
	shortMessage string) {
	pdu := deliverSM{}
	pdu.header.commandID = "deliver_sm"
	pdu.header.sequence = seq
	pdu.header.status = "ROK"

	if len(serviceType) <= 5 {
		pdu.serviceType = serviceType + "\x00"
	}

	pdu.sourceAddrTON = variables.TypeOfNumber[sourceAddrTON]
	pdu.sourceAddrNPI = variables.NumericPlanIndicator[sourceAddrNPI]

	if len(sourceAddr) <= 20 {
		pdu.sourceAddr = sourceAddr + "\x00"
	}

	pdu.destAddrTON = variables.TypeOfNumber[destAddrTON]
	pdu.destAddrNPI = variables.NumericPlanIndicator[destAddrNPI]

	if len(destinationAddr) <= 20 {
		pdu.destinationAddr = destinationAddr + "\x00"
	}

	if len(shortMessage) <= 254 {
		pdu.shortMessage = shortMessage
	}

}

func CreateDeliverSMRESP(seq uint32) deliverSMRESP {
	pdu := deliverSMRESP{}
	pdu.header.length = 17
	pdu.header.commandID = "deliver_sm_resp"
	pdu.header.status = "ROK"
	pdu.header.sequence = seq
	pdu.messageID = []byte{0}
	return pdu
}

func (pdu *deliverSMRESP) Pack() []byte {
	var buffer []byte
	buffer = append(buffer, pdu.header.Pack()...)
	buffer = append(buffer, pdu.messageID...)
	return buffer
}
