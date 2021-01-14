package _const

import (
	"encoding/json"
	"fmt"
	"net"
	"testing"
)

type AdapterNetConfig struct {
	Name       string           `json: Name`
	Index      int              `json: Index`
	MAC        net.HardwareAddr `json: Mac`
	IP         *net.IPNet       `json: IP`
	Gateway    string           `json: Gateway`
	DNSServers string           `json: DNSServers`
}

const PORT = 0xfffffffe

func deferTest() int {
	num := 1
	defer func() {
		fmt.Printf("Defer1, num: %v\n", num)
		if num = 2; num == 2 {
			fmt.Printf("Defer2, num: %v\n", num)
		}
	}()
	if num := 222; num > 1 {

	}
	return num
}

func PrintInt32(num uint32) {
	fmt.Print(num)
}

func TestConstNumberConvertion(t *testing.T) {
	PrintInt32(PORT & 0x0000ffff)
}

func TestConfigJson(t *testing.T) {
	macStr := "00:00:5e:00:53:01"
	mac, _ := net.ParseMAC(macStr)
	ipNet := net.IPNet{IP: net.ParseIP("192.168.0.2"), Mask: net.IPv4Mask(255, 255, 255, 0)}
	config := AdapterNetConfig{"Ethernet0", 1, mac, &ipNet, "192.168.1.2", "8.8.8.8"}
	configBytes, _ := json.Marshal(config)
	fmt.Print(string(configBytes))
	var config2 AdapterNetConfig
	json.Unmarshal(configBytes, &config2)
	fmt.Print(config2)
}

func TestDefer(t *testing.T) {
	num := deferTest()
	fmt.Print(num)
}
