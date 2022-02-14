package main

import (
	"fmt"
	"os/exec"
)

func main() {
	refresh_cmd := "sudo apt-get update"
	var _, _ = exec.Command("bash", "-c", refresh_cmd).Output()

	cmd := "sudo apt-get dist-upgrade -s | grep Inst"
	var out, _ = exec.Command("bash", "-c", cmd).Output()
	update_output := string(out)
	fmt.Println(update_output)
}
