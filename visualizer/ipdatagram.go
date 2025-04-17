package visualizer

import "fmt"

func IPDatagramVersion(buf []byte) (uint8, error) {
	if len(buf) < 20 {
		return 0, fmt.Errorf("buffer too small")
	}
	return buf[0] >> 4, nil
}

func IPDatagram(buf []byte) {
	// Check if the buffer is large enough to contain an IP header
	if len(buf) < 20 {
		return
	}
	// Version
	version, _ := IPDatagramVersion(buf)

	// Extract the IP header length (IHL) from the first byte
	ihl := int(buf[0] & 0x0F)

	// Calculate the total length of the IP header
	ipHeaderLength := ihl * 4

	// Check if the buffer is large enough to contain the entire IP datagram
	if len(buf) < ipHeaderLength {
		return
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

	// Print the source and destination IP addresses
	fmt.Printf("VERSION: %d\n", version)
	fmt.Printf("IHL: %d\n", ihl)
	fmt.Printf("Total Length: %d\n", totalLength)
	fmt.Printf("Identification: %d\n", identification)
	fmt.Printf("TTL: %d\n", buf[8])
	fmt.Printf("Protocol: %s\n", protocol)
	fmt.Printf("Header Checksum: %d\n", headerChecksum)
	fmt.Printf("SRC IP: %d.%d.%d.%d\n", int(buf[12]), int(buf[13]), int(buf[14]), int(buf[15]))
	fmt.Printf("DST IP: %d.%d.%d.%d\n", buf[16], buf[17], buf[18], buf[19])
}
