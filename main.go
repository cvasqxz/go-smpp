package main

import (
	"bufio"
	"log"
	"time"

	"github.com/cvasqxz/go-smpp/pdu"
	"github.com/cvasqxz/go-smpp/utils"
)

var sequence uint32

func SendEnquireLink(readwriter *bufio.ReadWriter) {
	for {
		time.Sleep(10 * time.Second)

		sequence++
		enquireLink := pdu.CreateEnquireLink("enquire_link", sequence)
		readwriter.Write(enquireLink.PackEnquireLink())
		readwriter.Flush()

		log.Println("SEND", sequence, "-> enquireLink")
	}
}

func GetResponses(readwriter *bufio.ReadWriter, channel chan []byte) {
	for {
		lenBuffer := make([]byte, 4)
		_, err := readwriter.Read(lenBuffer)
		utils.ErrorHandler(err)

		packetBuffer := make([]byte, utils.Byte2int(lenBuffer))
		_, err = readwriter.Read(packetBuffer)

		channel <- append(lenBuffer, packetBuffer...)
	}
}

func main() {
	channel := make(chan []byte)

	sequence = 0
	readwriter, err := utils.OpenConn("127.0.0.1:9661")
	utils.ErrorHandler(err)

	// send bind_receiver
	sequence++
	packet := pdu.CreateBind("receiver", sequence, "ID666", "CESAR", "Hola.123", "3.4", "unknown", "unknown")
	readwriter.Write(packet.PackBind())
	readwriter.Flush()
	log.Println("SEND", sequence, "-> bind_receiver")

	go SendEnquireLink(readwriter)
	go GetResponses(readwriter, channel)

	for {
		recv := <-channel
		recvHeader := pdu.ParseHeader(recv)
		sequence = recvHeader.GetSequence()
		log.Println("RECV", sequence, "<-", recvHeader.GetCommandID())
	}
}
