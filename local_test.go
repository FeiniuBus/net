package net

import (
	"net"
	"testing"
)

func TestIsPrivateSubnet(t *testing.T) {
	addrs := []string{
		"192.168.3.25",
		"172.16.31.34",
	}

	for _, addr := range addrs {
		ip := net.ParseIP(addr)
		result := IsPrivateSubnet(ip)

		if !result {
			t.Errorf("Expect true, actual is false, IP: %s", ip.String())
		}
	}
}
