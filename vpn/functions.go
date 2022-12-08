package vpn

import "ZoncordId/cmdCommands"

func StartVpn() {
	cmdCommands.Run(
		"wg-quick up zoncord",
	)
}

func StopVpn() {
	cmdCommands.Run(
		"wg-quick down zoncord",
	)
}
