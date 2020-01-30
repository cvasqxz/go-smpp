package main

import (
	"log"

	"github.com/cvasqxz/go-smpp/structs"
)

func main() {
	packet := structs.CreateBind("transmitter", 1, "cesar", "Hola.123", "notebook", "national", "wap")
	log.Println(packet.PackBind())
}
