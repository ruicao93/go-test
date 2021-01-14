package main

import (
	"flag"
	"fmt"
	"github.com/Microsoft/hcsshim"
)

var (
	netName = flag.String("net", "antrea-hnsnetwork", "hnsnetwork name")
)

func main() {
	net, err := hcsshim.GetHNSNetworkByName(*netName)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Net: %v\n", net)
}
