package net

import "net"

// InterfaceIPV4Addrs returns all ip v4 addrs
func InterfaceIPV4Addrs() ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	result := make([]net.IP, 0, len(addrs))
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip := ipnet.IP.To4(); ip != nil {
				result = append(result, ip)
			}
		}
	}

	return result, nil
}
