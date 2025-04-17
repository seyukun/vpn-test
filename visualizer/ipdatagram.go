package visualizer

import (
	"fmt"
	"net"
)

func IPDatagramVersion(buf []byte) (uint8, error) {
	if len(buf) < 20 {
		return 0, fmt.Errorf("buffer too small")
	}
	return buf[0] >> 4, nil
}

func IPDatagramV4(buf []byte) (string, error) {
	// Version
	version, err := IPDatagramVersion(buf)
	if err != nil {
		return "", err
	}

	// Extract the IP header length (IHL) from the first byte
	ihl := int(buf[0] & 0x0F)

	// Calculate the total length of the IP header
	ipHeaderLength := ihl * 4

	// Check if the buffer is large enough to contain the entire IP datagram
	if len(buf) < ipHeaderLength {
		return "", fmt.Errorf("header too small")
	}

	// Total Length
	totalLength := int(buf[2])<<8 | int(buf[3])

	// Identification
	identification := int(buf[4])<<8 | int(buf[5])

	// Header Checksum
	headerChecksum := int(buf[10])<<8 | int(buf[11])

	// Protocol
	protocolNumber := int(buf[9])
	protocol := "Unknown"
	switch protocolNumber {
	case 0:
		protocol = "Reserved"
	case 1:
		protocol = "ICMP"
	case 2:
		protocol = "IGMP"
	case 6:
		protocol = "TCP"
	case 8:
		protocol = "EGP"
	case 17:
		protocol = "UDP"
	case 50:
		protocol = "ESP"
	case 51:
		protocol = "AH"
	}

	output := fmt.Sprintf("┌────┬────┬─────────┬───────────────────┐\n")
	output += fmt.Sprintf("│v%-3d│%-4d│%08b │%-19d│\n", version, ihl, buf[1], totalLength)
	output += fmt.Sprintf("├────┴────┴─────────┼───────────────────┤\n")
	output += fmt.Sprintf("│%-19d│%016b   │\n", identification, uint16(buf[6])<<8|uint16(buf[7]))
	output += fmt.Sprintf("├─────────┬─────────┼───────────────────┤\n")
	output += fmt.Sprintf("│%08b │%-9s│%-19d│\n", buf[8], protocol, headerChecksum)
	output += fmt.Sprintf("├─────────┴─────────┴───────────────────┤\n")
	output += fmt.Sprintf("│%-39s│\n", fmt.Sprintf("%d.%d.%d.%d", buf[12], buf[13], buf[14], buf[15]))
	output += fmt.Sprintf("├───────────────────────────────────────┤\n")
	output += fmt.Sprintf("│%-39s│\n", fmt.Sprintf("%d.%d.%d.%d", buf[16], buf[17], buf[18], buf[19]))
	output += fmt.Sprintf("└───────────────────────────────────────┘\n")
	return output, nil
}

func IPDatagramV6(buf []byte) (string, error) {
	// Version
	version, err := IPDatagramVersion(buf)
	if err != nil {
		return "", err
	}

	// Extract the Traffic Class and Flow Label
	trafficClass := (buf[0]&0b00001111)<<4 | (buf[1]&0b11110000)>>4
	flowLavel := (uint32(buf[1]&0b00001111)<<16 | uint32(buf[2])<<8 | uint32(buf[3]))

	// Payload Length
	payloadLength := uint16(buf[4])<<8 | uint16(buf[5])

	// Next Header
	nextHaederNumber := int(buf[6])
	nextHaeder := "Unknown"
	switch nextHaederNumber {
	case 0:
		nextHaeder = "Reserved"
	case 1:
		nextHaeder = "ICMP"
	case 2:
		nextHaeder = "IGMP"
	case 6:
		nextHaeder = "TCP"
	case 8:
		nextHaeder = "EGP"
	case 17:
		nextHaeder = "UDP"
	case 50:
		nextHaeder = "ESP"
	case 51:
		nextHaeder = "AH"
	}

	// Hop Limit
	hopLimit := int(buf[7])

	output := fmt.Sprintf("┌────┬─────────┬────────────────────────┐\n")
	output += fmt.Sprintf("│v%-3d│%08b │%020d    │\n", version, trafficClass, flowLavel)
	output += fmt.Sprintf("├────┴─────────┴────┬─────────┬─────────┤\n")
	output += fmt.Sprintf("│%-19d│%-9s|%-9d│\n", payloadLength, nextHaeder, hopLimit)
	output += fmt.Sprintf("├───────────────────┴─────────┴─────────┤\n")
	output += fmt.Sprintf("│%-39s│\n", net.IP(buf[8:24]))
	output += fmt.Sprintf("├───────────────────────────────────────┤\n")
	output += fmt.Sprintf("│%-39s│\n", net.IP(buf[24:40]))
	output += fmt.Sprintf("└───────────────────────────────────────┘\n")

	return output, nil
}

func IPDatagram(buf []byte) {
	// Version
	version, err := IPDatagramVersion(buf)
	if err != nil {
		return
	}

	switch version {
	case 4:
		output, err := IPDatagramV4(buf)
		if err != nil {
			return
		}
		fmt.Println(output)

	case 6:
		if len(buf) < 40 {
			fmt.Println("Header too small for IPv6")
			return
		}
		output, err := IPDatagramV6(buf)
		if err != nil {
			return
		}
		fmt.Println(output)

	default:
		fmt.Println("Unsupported IP version:", version)
	}
}

func IPDatagramTUN(buf []byte) {
	IPDatagram(buf[4:])
}
