package main

import (
	"flag"
	"github.com/vmware-tanzu/antrea/pkg/ovs/ovsconfig"
	"k8s.io/klog"
)

var (
	bridgeName      = flag.String("br", "vEthernet (Ethernet0 2)", "Bridge name")
	bridgeInterface = flag.String("interface", "vEthernet (Ethernet0 2)", "Interface name")
)

const OVSRUNDir = `C:\openvswitch\var\run\openvswitch`

func main() {
	ovsdbAddress := ovsconfig.GetConnAddress(OVSRUNDir)
	ovsdbConnection, err := ovsconfig.NewOVSDBConnectionUDS(ovsdbAddress)
	if err != nil {
		klog.Errorf("error connecting OVSDB: %v", err)
		return
	}
	defer ovsdbConnection.Close()

	ovsBridgeClient := ovsconfig.NewOVSBridge(*bridgeName, "system", ovsdbConnection)

	if err = ovsBridgeClient.Create(); err != nil {
		klog.Errorf("Failed to create bridge: %v", err)
		return
	}
	_, err = ovsBridgeClient.GetOFPort(*bridgeInterface)
	if err == nil {
		klog.Infof("Bridge interface already exists.")
		return
	}
	if _, err = ovsBridgeClient.CreateInternalPort(*bridgeName, 0, nil); err != nil {
		klog.Errorf("Failed to add bridge interface: %v", err)
	}
}
