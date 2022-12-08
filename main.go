package main

import (
	"ZoncordId/conf"
	"ZoncordId/vpn"
	"time"
)

func main() {
	conf.GenerateConfig()
	vpn.StartVpn()
	time.Sleep(100)
	vpn.StopVpn()
}
