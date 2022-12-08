package main

import (
	"ZoncordId/conf"
	"ZoncordId/vpn"
	flag2 "flag"
)

func main() {
	conf.GenerateConfig()
	startVpn := flag2.Bool("startVpn", false, "start vpn")
	stopVpn := flag2.Bool("stopVpn", false, "stop vpn")
	flag2.Parse()

	if *startVpn {
		vpn.StartVpn()
	}
	if *stopVpn {
		vpn.StopVpn()
	}

}
