package visualizer

import "fmt"

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
	default:
		fmt.Println("Unsupported IP version:", version)
	}
}
