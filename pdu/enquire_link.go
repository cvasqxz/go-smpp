package pdu

type enquireLink struct {
	header Header
}

func CreateEnquireLink(enquire_type string, seq uint32) enquireLink {
	pdu := enquireLink{}
	pdu.header.length = 16
	pdu.header.commandID = enquire_type
	pdu.header.status = "ROK"
	pdu.header.sequence = seq
	return pdu
}

func (pdu *enquireLink) Pack() []byte {
	return pdu.header.Pack()
}
