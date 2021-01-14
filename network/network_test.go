package network

import (
	"fmt"
	"gotest.tools/assert"
	"net"
	"strings"
	"testing"

	"github.com/ruicao93/go-test/utils"
)

func TestDupAddr(t *testing.T) {
	// Instance MSFT_NetIPAddress already exists
	ip, ipnet, err := net.ParseCIDR("192.168.31.68/24")
	assert.NilError(t, err)
	testIPNet := net.IPNet{IP: ip, Mask: ipnet.Mask}

	t.Logf("test ip: %s", testIPNet.String())
	err = utils.ConfigureInterfaceAddressWithDefaultGateway("WLAN 2", &testIPNet)
	if err != nil {
		t.Logf("error f: %v", err)
		if strings.Contains(err.Error(), "Instance MSFT_NetIPAddress already exists") {
			t.Logf("address already exists")
		}
	}
}

func TestListInterfaces(t *testing.T) {
	inf, _ := net.InterfaceByName("WLAN 2")
	addrs, _ := inf.Addrs()
	for _, addr := range addrs {
		t.Logf("%s", addr.String())
	}
}

func TestError(t *testing.T) {
	err := fmt.Errorf("original error")
	if inf, err := net.InterfaceByName("heheda"); err != nil {
		t.Logf("%v: %s", inf, err.Error())
	}
	t.Logf("%s", err.Error())
}
