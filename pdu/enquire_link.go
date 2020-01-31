package pdu

type enquireLink struct {
	header Header
}

func CreateEnquireLink(seq uint32) enquireLink {
	pdu := enquireLink{}
	pdu.header.length = 16
	pdu.header.commandID = "enquire_link"
	pdu.header.status = "ROK"
	pdu.header.sequence = seq
	return pdu
}

func (pdu *enquireLink) Pack() []byte {
	return pdu.header.Pack()
}

type enquireLinkRESP struct {
	header Header
}

func CreateEnquireLinkRESP(seq uint32) enquireLinkRESP {
	pdu := enquireLinkRESP{}
	pdu.header.length = 16
	pdu.header.commandID = "enquire_link_resp"
	pdu.header.status = "ROK"
	pdu.header.sequence = seq
	return pdu
}

func (pdu *enquireLinkRESP) Pack() []byte {
	return pdu.header.Pack()
}
