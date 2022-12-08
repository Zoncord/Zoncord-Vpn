package cmdCommands

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Run(command string) {
	slices := strings.Split(command, " ")
	fmt.Println(slices[0], slices[1:])
	cmd := exec.Command(slices[0], slices[1:]...)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return
	}
}
func RunOutput(command string) {
	slices := strings.Split(command, " ")
	cmd := exec.Command(slices[0], slices[0:]...)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(output)
}

func RunMany(commands ...string) {
	for i := 0; i < len(commands); i++ {
		Run(commands[i])
	}
}
