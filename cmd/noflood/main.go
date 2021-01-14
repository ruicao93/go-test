package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
)

var ofport = flag.Int("ofport", 6262, "Input ofport")

func main() {
	flag.Parse()
	cmdStr := fmt.Sprintf("ovs-ofctl mod-port br-int %d no-flood", *ofport)
	cmd := exec.Command("cmd.exe", "/c", cmdStr)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s", string(stderr.Bytes()))
	}
}
