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

	// AVPs to add
	start.AVPS = make([]goradius.AttributeValuePair, 0, 10)

	start.AVPS = append(start.AVPS, goradius.AttributeValuePair{"User-Name", // name
		"string",   // type
		7,          // length
		"john doe", // content
		1})         // avp type

	// Acct-Session-Id
	start.AVPS = append(start.AVPS, goradius.AttributeValuePair{"Acct-Session-Id", // name
		"string", // type
		1,        // length
		"3",      // content
		44})      // avp type

	// Framed-IP-Address (8)
	start.AVPS = append(start.AVPS, goradius.AttributeValuePair{"Framed-IP-Address", // name
		"IPAddress", // type
		4,           // length
		"1.2.3.4",   // content
		8})          // avp type

	// Calling-Station-Id (31)
	start.AVPS = append(start.AVPS, goradius.AttributeValuePair{"Calling-Station-Id", // name
		"string",      // type
		10,            // length
		"39955555528", // content
		31})           // avp type

	// NAS-Identifer(32)
	start.AVPS = append(start.AVPS, goradius.AttributeValuePair{"NAS-Identifier", // name
		"string",      // type
		11,            // length
		"10.10.10.10", // content
		32})           // avp type

	// NAS-IP-Address(4)
	start.AVPS = append(start.AVPS, goradius.AttributeValuePair{"NAS-IP-Address", // name
		"IPAddress",   // type
		4,             // length
		"10.10.10.10", // content
		4})            // avp type

	fmt.Printf("Radius packet: %s\n", start)
	r.SendPacketToServer(start)
}

func SampleSharedSecret(nas string) string {
	return "testing123"
}
