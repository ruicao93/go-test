package route

import (
	"fmt"
	"testing"

	"github.com/rakelkar/gonetsh/netroute"
	"github.com/stretchr/testify/assert"
)

func TestListRoutes(t *testing.T) {
	nr := netroute.New()
	routes, err := nr.GetNetRoutesAll()
	ifIndex := 20
	gwIP := "192.168.31.1"

	assert.Equal(t, err, nil)
	for idx := range routes {
		route := routes[idx]
		if route.LinkIndex != ifIndex {
			continue
		}
		if route.DestinationSubnet.IP.To4() == nil {
			continue
		}
		fmt.Printf("Test route: %v\n", route)
		if route.GatewayAddress.String() != gwIP {
			continue
		}
		if route.DestinationSubnet.IP.IsUnspecified() {
			continue
		}
		fmt.Printf("route: %v\n", route)
	}
}
