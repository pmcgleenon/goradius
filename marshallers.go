package goradius

import (
	"net"
)

func StringMarshaller(a AttributeValuePair, p *RadiusPacket) []byte {

	avpLen := (byte)(2 + len(a.Content.(string)))

	avp := []byte{a.AVPType, avpLen}

	value := []byte(a.Content.(string))

	avp = append(avp, value...)

	return avp

}

func AcctStatusTypeMarshaller(a AttributeValuePair, p *RadiusPacket) []byte {

	var avpLen byte
	avpLen = 3

	avp := []byte{a.AVPType, avpLen}

	var acctStatus byte
	switch a.Content.(string) {
	case "Start":
		acctStatus = 1
	case "Stop":
		acctStatus = 2
	case "Interim-Update":
		acctStatus = 3
	case "Accounting-On":
		acctStatus = 4
	case "Accounting-Off":
		acctStatus = 5
	}

	avp = append(avp, acctStatus)
	return avp
}

func IPAddressMarshaller(a AttributeValuePair, p *RadiusPacket) []byte {

	var avpLen byte
	avpLen = 6

	ip4 := net.ParseIP(a.Content.(string))

	avp := []byte{a.AVPType, avpLen}

	avp = append(avp, ip4.To4()...)
	return avp
}
