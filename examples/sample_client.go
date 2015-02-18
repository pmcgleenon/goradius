package main

import (
	"fmt"
	"github.com/pmcgleenon/goradius"
)

func main() {
	r := goradius.NewGoRadius("../dict/radius-dict.json",
		"../dict/vendor-dict.json",
		true, /* debug */
		true /* verbose */)

	if r == nil {
		fmt.Printf("Error initializing Radius: %s\n")
		return
	}

	err := r.DialUDPServer("10.20.51.12:1813")

	if err != nil {
		fmt.Printf("Unable to create UDP socket : %s\n", err.Error())
		return
	}

	start := new(goradius.RadiusPacket)
	start.Code = uint(4)
	start.PacketId = 12345
	//    start.Authenticator = null
	start.SharedSecret = "testing123"

	start.AVPS = make([]goradius.AttributeValuePair, 0, 10)
	start.AVPS = append(start.AVPS, goradius.AttributeValuePair{"User-Name", // name
		"string",  // type
		7,         // length
		"john doe", // content
		1})        // avp type
	fmt.Printf("Radius packet: %s\n", start)

	r.SendPacketToServer(start)
}

func SampleSharedSecret(nas string) string {
	return "testing123"
}
