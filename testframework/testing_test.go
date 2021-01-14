package testframework

import "testing"

var hypervInstalled = false

func TestWindowsHyperVInstalled(t *testing.T) {
	hypervInstalled = true
}

func TestCreateHnsNetwork(t *testing.T) {
	t.Logf("HyperV installation state: %t", hypervInstalled)
}
