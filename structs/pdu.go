package structs

type replaceSmResp struct {
	header Header
}

type genericNACK struct {
	header Header
}

type cancelSmRESP struct {
	header Header
}

type submitSm struct {
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
	scheduleDeliveryTime string
	validityPeriod       string
	registeredDelivery   []byte
	replaceIfPresentFlag []byte
	dataCoding           []byte
	smDefaultMsgID       []byte
	smLength             []byte
	shortMessage         string
}

type submitSmRESP struct {
	header    Header
	messageID string
}

type submitMulti struct {
	header               Header
	serviceType          string
	sourceAddrTON        []byte
	sourceAddrNPI        []byte
	sourceAddr           string
	numberOfDest         []byte
	destAddress          string
	esmClass             []byte
	protocolID           []byte
	priorityFlag         []byte
	scheduleDeliveryTime string
	validityPeriod       string
	registeredDelivery   []byte
	replaceIfPresentFlag []byte
	dataCoding           []byte
	smDefaultMsgID       []byte
	smLength             []byte
	shortMessage         string
}

type destAddress struct {
	header   Header
	destFlag []byte
}

type smeDestAddress struct {
	header          Header
	destAddrTON     []byte
	destAddrNPI     []byte
	destinationAddr string
}

type distributionList struct {
	header Header
	dlName string
}

type submitMultiRESP struct {
	header       Header
	messageID    string
	noUnsuccess  []byte
	unsuccessSme []byte
}

type unsuccessSme struct {
	header          Header
	destAddrTON     []byte
	destAddrNPI     []byte
	destinationAddr string
	errorStatusCode []byte
}

type deliverSm struct {
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
	scheduleDeliveryTime string
	validityPeriod       string
	registeredDelivery   []byte
	replaceIfPresentFlag []byte
	dataCoding           []byte
	smDefaultMsgID       []byte
	smLength             []byte
	shortMessage         string
}

type deliverSmRESP struct {
	header    Header
	messageID string
}

type dataSm struct {
	header             Header
	serviceType        string
	sourceAddrTON      []byte
	sourceAddrNPI      []byte
	sourceAddr         string
	destAddrTON        []byte
	destAddrNPI        []byte
	destinationAddr    string
	esmClass           []byte
	registeredDelivery []byte
	dataCoding         []byte
}

type dataSmRESP struct {
	header    Header
	messageID string
}

type querySm struct {
	header        Header
	messageID     string
	sourceAddrTON []byte
	sourceAddrNPI []byte
	sourceAddr    string
}

type querySmRESP struct {
	header       Header
	messageID    string
	finalDate    string
	messageState []byte
	errorCode    []byte
}

type cancelSm struct {
	header          Header
	serviceType     string
	messageID       string
	sourceAddrTON   []byte
	sourceAddrNPI   []byte
	sourceAddr      string
	destAddrTON     []byte
	destAddrNPI     []byte
	destinationAddr string
}

type replaceSm struct {
	header               Header
	messageID            string
	sourceAddrTON        []byte
	sourceAddrNPI        []byte
	sourceAddr           string
	scheduleDeliveryTime string
	validityPeriod       string
	registeredDelivery   []byte
	replaceIfPresentFlag []byte
	dataCoding           []byte
	smDefaultMsgID       []byte
	smLength             []byte
	shortMessage         string
}

type alertNotification struct {
	header        Header
	sourceAddrTON []byte
	sourceAddrNPI []byte
	sourceAddr    string
	esmeAddrTON   []byte
	esmeAddrNPI   []byte
	esmeAddr      string
}
