package netx

import "net"

// InternalIp returns an internal ip.
func InternalIp() string {
	infs, err := net.Interfaces()
	if err != nil {
		return ""
	}

	var ip string
	for _, inf := range infs {
		if isEthDown(inf.Flags) || isLoopback(inf.Flags) {
			continue
		}

		addrs, err := inf.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					// return ipnet.IP.String()
					tmp := ipnet.IP.String()
					if tmp[:3] == "192" {
						// 优先使用192网段IP
						return tmp
					} else if len(ip) == 0 {
						ip = tmp
					}
				}
			}
		}
	}

	return ip
}

func isEthDown(f net.Flags) bool {
	return f&net.FlagUp != net.FlagUp
}

func isLoopback(f net.Flags) bool {
	return f&net.FlagLoopback == net.FlagLoopback
}
