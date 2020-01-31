package main

import (
	"log"

	"github.com/cvasqxz/go-smpp/pdu"
	"github.com/cvasqxz/go-smpp/utils"
)

func main() {
	readwriter, err := utils.OpenConn("127.0.0.1:9661")
	utils.ErrorHandler(err)

	// send bind_receiver
	packet := pdu.CreateBind("transmitter", 0, "ID666", "CESAR", "Hola.123", "3.4", "unknown", "unknown")
	readwriter.Write(packet.PackBind())
	readwriter.Flush()

	// read bind_receiver_resp
	lenBuffer := make([]byte, 4)
	_, err = readwriter.Read(lenBuffer)
	utils.ErrorHandler(err)

	packetBuffer := make([]byte, utils.Byte2int(lenBuffer))
	_, err = readwriter.Read(packetBuffer)

	packetRECV := append(lenBuffer, packetBuffer...)

	log.Println("RECV: ", pdu.ParseHeader(packetRECV))
}
