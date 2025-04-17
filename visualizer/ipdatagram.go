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

func IPDatagramProtocol(protocolNumber uint8) (string, error) {
	protocol := fmt.Sprintf("??(%d)", protocolNumber)
	switch protocolNumber {
	case 0:
		protocol = "HOPOPT"
	case 1:
		protocol = "ICMP"
	case 2:
		protocol = "IGMP"
	case 3:
		protocol = "GGP"
	case 4:
		protocol = "IPv4"
	case 5:
		protocol = "ST"
	case 6:
		protocol = "TCP"
	case 7:
		protocol = "CBT"
	case 8:
		protocol = "EGP"
	case 9:
		protocol = "IGP"
	case 10:
		protocol = "BBN-RCC-MON"
	case 11:
		protocol = "NVP-II"
	case 12:
		protocol = "PUP"
	case 13:
		protocol = "ARGUS (deprecated)"
	case 14:
		protocol = "EMCON"
	case 15:
		protocol = "XNET"
	case 16:
		protocol = "CHAOS"
	case 17:
		protocol = "UDP"
	case 18:
		protocol = "MUX"
	case 19:
		protocol = "DCN-MEAS"
	case 20:
		protocol = "HMP"
	case 21:
		protocol = "PRM"
	case 22:
		protocol = "XNS-IDP"
	case 23:
		protocol = "TRUNK-1"
	case 24:
		protocol = "TRUNK-2"
	case 25:
		protocol = "LEAF-1"
	case 26:
		protocol = "LEAF-2"
	case 27:
		protocol = "RDP"
	case 28:
		protocol = "IRTP"
	case 29:
		protocol = "ISO-TP4"
	case 30:
		protocol = "NETBLT"
	case 31:
		protocol = "MFE-NSP"
	case 32:
		protocol = "MERIT-INP"
	case 33:
		protocol = "DCCP"
	case 34:
		protocol = "3PC"
	case 35:
		protocol = "IDPR"
	case 36:
		protocol = "XTP"
	case 37:
		protocol = "DDP"
	case 38:
		protocol = "IDPR-CMTP"
	case 39:
		protocol = "TP++"
	case 40:
		protocol = "IL"
	case 41:
		protocol = "IPv6"
	case 42:
		protocol = "SDRP"
	case 43:
		protocol = "IPv6-Route"
	case 44:
		protocol = "IPv6-Frag"
	case 45:
		protocol = "IDRP"
	case 46:
		protocol = "RSVP"
	case 47:
		protocol = "GRE"
	case 48:
		protocol = "DSR"
	case 49:
		protocol = "BNA"
	case 50:
		protocol = "ESP"
	case 51:
		protocol = "AH"
	case 52:
		protocol = "I-NLSP"
	case 53:
		protocol = "SWIPE (deprecated)"
	case 54:
		protocol = "NARP"
	case 55:
		protocol = "Min-IPv4"
	case 56:
		protocol = "TLSP"
	case 57:
		protocol = "SKIP"
	case 58:
		protocol = "IPv6-ICMP"
	case 59:
		protocol = "IPv6-NoNxt"
	case 60:
		protocol = "IPv6-Opts"
	case 62:
		protocol = "CFTP"
	case 64:
		protocol = "SAT-EXPAK"
	case 65:
		protocol = "KRYPTOLAN"
	case 66:
		protocol = "RVD"
	case 67:
		protocol = "IPPC"
	case 69:
		protocol = "SAT-MON"
	case 70:
		protocol = "VISA"
	case 71:
		protocol = "IPCV"
	case 72:
		protocol = "CPNX"
	case 73:
		protocol = "CPHB"
	case 74:
		protocol = "WSN"
	case 75:
		protocol = "PVP"
	case 76:
		protocol = "BR-SAT-MON"
	case 77:
		protocol = "SUN-ND"
	case 78:
		protocol = "WB-MON"
	case 79:
		protocol = "WB-EXPAK"
	case 80:
		protocol = "ISO-IP"
	case 81:
		protocol = "VMTP"
	case 82:
		protocol = "SECURE-VMTP"
	case 83:
		protocol = "VINES"
	case 84:
		protocol = "IPTM"
	case 85:
		protocol = "NSFNET-IGP"
	case 86:
		protocol = "DGP"
	case 87:
		protocol = "TCF"
	case 88:
		protocol = "EIGRP"
	case 89:
		protocol = "OSPFIGP"
	case 90:
		protocol = "Sprite-RPC"
	case 91:
		protocol = "LARP"
	case 92:
		protocol = "MTP"
	case 93:
		protocol = "AX.25"
	case 94:
		protocol = "IPIP"
	case 95:
		protocol = "MICP (deprecated)"
	case 96:
		protocol = "SCC-SP"
	case 97:
		protocol = "ETHERIP"
	case 98:
		protocol = "ENCAP"
	case 100:
		protocol = "GMTP"
	case 101:
		protocol = "IFMP"
	case 102:
		protocol = "PNNI"
	case 103:
		protocol = "PIM"
	case 104:
		protocol = "ARIS"
	case 105:
		protocol = "SCPS"
	case 106:
		protocol = "QNX"
	case 107:
		protocol = "A/N"
	case 108:
		protocol = "IPComp"
	case 109:
		protocol = "SNP"
	case 110:
		protocol = "Compaq-Peer"
	case 111:
		protocol = "IPX-in-IP"
	case 112:
		protocol = "VRRP"
	case 113:
		protocol = "PGM"
	case 115:
		protocol = "L2TP"
	case 116:
		protocol = "DDX"
	case 117:
		protocol = "IATP"
	case 118:
		protocol = "STP"
	case 119:
		protocol = "SRP"
	case 120:
		protocol = "UTI"
	case 121:
		protocol = "SMP"
	case 122:
		protocol = "SM (deprecated)"
	case 123:
		protocol = "PTP"
	case 124:
		protocol = "ISIS over IPv4"
	case 125:
		protocol = "FIRE"
	case 126:
		protocol = "CRTP"
	case 127:
		protocol = "CRUDP"
	case 128:
		protocol = "SSCOPMCE"
	case 129:
		protocol = "IPLT"
	case 130:
		protocol = "SPS"
	case 131:
		protocol = "PIPE"
	case 132:
		protocol = "SCTP"
	case 133:
		protocol = "FC"
	case 134:
		protocol = "RSVP-E2E-IGNORE"
	case 135:
		protocol = "Mobility Header"
	case 136:
		protocol = "UDPLite"
	case 137:
		protocol = "MPLS-in-IP"
	case 138:
		protocol = "manet"
	case 139:
		protocol = "HIP"
	case 140:
		protocol = "Shim6"
	case 141:
		protocol = "WESP"
	case 142:
		protocol = "ROHC"
	case 143:
		protocol = "Ethernet"
	case 144:
		protocol = "AGGFRAG"
	case 145:
		protocol = "NSH"
	case 146:
		protocol = "Homa"
	case 147:
		protocol = "BIT-EMU"
	case 255:
		protocol = "Reserved"
	default:
		return protocol, fmt.Errorf("unknown protocol number: %d", protocolNumber)
	}
	return protocol, nil
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
	protocol, _ := IPDatagramProtocol(buf[9])

	output := fmt.Sprintf(`
\r┌────┬────┬─────────┬───────────────────┐
│v%-3d│%-4d│%08b │%-19d│
├────┴────┴─────────┼───────────────────┤
│%-19d│%016b   │
├─────────┬─────────┼───────────────────┤
│%08b │%-9s│%-19d│
├─────────┴─────────┴───────────────────┤
│%-39s│
├───────────────────────────────────────┤
│%-39s│
└───────────────────────────────────────┘`, version, ihl, buf[1], totalLength, identification, uint16(buf[6])<<8|uint16(buf[7]), buf[8], protocol, headerChecksum, fmt.Sprintf("%d.%d.%d.%d", buf[12], buf[13], buf[14], buf[15]), fmt.Sprintf("%d.%d.%d.%d", buf[16], buf[17], buf[18], buf[19]))

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
	nextHeader, _ := IPDatagramProtocol(buf[6])

	// Hop Limit
	hopLimit := int(buf[7])

	output := fmt.Sprintf(`
\r┌────┬─────────┬────────────────────────┐
│v%-3d│%08b │%020d    │
├────┴─────────┴────┬─────────┬─────────┤
│%-19d│%-9s|%-9d│
├───────────────────┴─────────┴─────────┤
│%-39s│
├───────────────────────────────────────┤
│%-39s│
└───────────────────────────────────────┘`, version, trafficClass, flowLavel, payloadLength, nextHeader, hopLimit, net.IP(buf[8:24]), net.IP(buf[24:40]))

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
