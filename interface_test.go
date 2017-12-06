package net

import (
	"testing"
)

func TestInterfaceIPV4Addrs(t *testing.T) {
	ips, err := InterfaceIPV4Addrs()
	if err != nil {
		t.Error(err)
	}

	for _, ip := range ips {
		if !IsPrivateSubnet(ip) {
			t.Errorf("Expect IPv4, IP: %s", ip.String())
		}

		t.Log(ip.String())
	}
}
