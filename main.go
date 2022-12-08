package main

import (
	"ZoncordId/conf"
	"ZoncordId/vpn"
	flag2 "flag"
)

func main() {
	conf.GenerateConfig()
	startVpn := flag2.Bool("start", false, "start vpn")
	stopVpn := flag2.Bool("stop", false, "stop vpn")
	flag2.Parse()

	if *startVpn {
		vpn.StartVpn()
	}
	if *stopVpn {
		vpn.StopVpn()
	}
}
