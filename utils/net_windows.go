package utils

import (
	"fmt"
	ps "github.com/benmoss/go-powershell"
	"github.com/benmoss/go-powershell/backend"
	"net"
	"strings"
)

func CallPSCommand(cmd string) (string, error) {
	// Create a backend shell.
	back := &backend.Local{}

	// start a local powershell process
	shell, err := ps.New(back)
	if err != nil {
		return "", err
	}
	defer shell.Exit()
	stdout, stderr, err := shell.Execute(cmd)
	if err != nil {
		return stdout, err
	}
	if stderr != "" {
		return stdout, fmt.Errorf("%s", stderr)
	}
	return stdout, nil
}

func InvokePSCommand(cmd string) error {
	_, err := CallPSCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}

func ConfigureInterfaceAddressWithDefaultGateway(ifaceName string, ipConfig *net.IPNet) error {
	ipStr := strings.Split(ipConfig.String(), "/")
	cmd := fmt.Sprintf(`New-NetIPAddress -InterfaceAlias "%s" -IPAddress %s -PrefixLength %s`, ifaceName, ipStr[0], ipStr[1])
	return InvokePSCommand(cmd)
}
