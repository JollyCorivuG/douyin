package utils

import (
	"net"
)

// func GetLocalIp() string {
// 	addrs, err := net.InterfaceAddrs()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	for _, value := range addrs {
// 		if ipNet, ok := value.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
// 			if ipNet.IP.To4() != nil {
// 				return ipNet.IP.String()
// 			}
// 		}
// 	}
// 	return ""
// }

func LocalIPv4() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String()
		}
	}
	return ""
}
