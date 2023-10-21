package vpn

import (
	"ZoncordId/cmdCommands"
	"fmt"
)

func StartVpn() {
	cmdCommands.Run(
		"wg-quick up zoncord",
	)
	fmt.Println("Vpn started")
}

func StopVpn() {
	cmdCommands.Run(
		"wg-quick down zoncord",
	)
	fmt.Println("Vpn stopped")
}
