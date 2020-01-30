package main

import (
	"log"

	"github.com/cvasqxz/go-smpp/structs"
)

func main() {
	packet := structs.CreateBind("transmitter", 30, "cesar", "Hola.123", "notebook", "3.4", "unknown", "unknown")
	log.Println(packet.PackBind())
	packet = structs.CreateBind("receiver", 30, "cesar", "Hola.123", "notebook", "3.4", "unknown", "unknown")
	log.Println(packet.PackBind())
	
}
