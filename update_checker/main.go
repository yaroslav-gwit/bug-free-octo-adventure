package main

import (
	"fmt"
	"os/exec"
)

func main() {
	refresh_cmd := "sudo apt-get update"
	var _, _ = exec.Command("bash", "-c", refresh_cmd).Output()

	all_updates_cmd := "sudo apt-get dist-upgrade -s | grep Inst"
	var all_updates_out, _ = exec.Command("bash", "-c", all_updates_cmd).Output()
	all_updates_output := string(all_updates_out)
	fmt.Println(all_updates_output)

	security_updates_cmd := "sudo apt-get dist-upgrade -s | grep Inst | grep security"
	var security_updates_out, _ = exec.Command("bash", "-c", security_updates_cmd).Output()
	security_updates_output := string(security_updates_out)
	fmt.Println(security_updates_output)
}
