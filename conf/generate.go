package conf

import (
	"ZoncordId/utils"
	"fmt"
	"log"
	"os"
	"runtime"
)

func writeConf(data []byte) {
	err := os.WriteFile("/etc/wireguard/zoncord.conf", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateConfig() {
	data := fmt.Sprintf(
		"[Interface]\nPrivateKey=%s\nAddress=%s\nDNS=%s\n[Peer]\n"+
			"PublicKey=%s\nAllowedIPs=0.0.0.0/0\nEndpoint=%s\nPersistentKeepalive=20",
		utils.GetEnv("PrivateKey"),
		utils.GetEnv("Address"),
		utils.GetEnv("DNS"),
		utils.GetEnv("PublicKey"),
		utils.GetEnv("Endpoint"),
	)
	bytes := []byte(data)

	switch system := runtime.GOOS; system {
	case "darwin":
		generateMac(bytes)
	case "linux":
		generateLinux(bytes)
	case "windows":
		generateWindows(bytes)
	}
}

func generateLinux(data []byte) {
	if _, err := os.Stat("/etc/wireguard/zoncord.conf"); err == nil {
		fileContent, err := os.ReadFile("/etc/wireguard/zoncord.conf")
		if err != nil {
			log.Fatal(err)
		}
		if len(fileContent) == len(data) {
			for i := 0; i < len(fileContent); i++ {
				if fileContent[i] != data[i] {
					writeConf(data)
					break
				}
			}
		} else {
			writeConf(data)
		}
	} else {
		writeConf(data)
	}
}

func generateWindows(data []byte) {

}

func generateMac(data []byte) {

}
