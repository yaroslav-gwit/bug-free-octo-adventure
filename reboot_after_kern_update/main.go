package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/facette/natsort"
)

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"

func main() {
	var running_kernel = cmdUname()
	var latest_kernel = cmdLsBoot()

	if running_kernel == latest_kernel {
		fmt.Println(colorGreen + "All good. You are running the latest kernel version." + colorReset)
	} else {
		fmt.Println(colorRed + "Please reboot to apply the kernel update!" + colorReset)
		fmt.Println()
		fmt.Println("You are running:      " + colorRed + running_kernel + colorReset)
		fmt.Println("The latest installed: " + colorGreen + latest_kernel + colorReset)
	}
}

func cmdUname() string {
	var cmd = "uname -r"
	var out, err = exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	var final_output = strings.ReplaceAll(string(out), "\n", "")
	return final_output
}

func cmdLsBoot() string {
	var cmd = "ls -1 /boot/ | grep \"vmlinuz\\|vmlinux\" | grep -v \"vmlinuz.old\\|vmlinux.old\""
	var out, err = exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	var output_ls = string(out)
	var output_ls_slices = strings.Split(output_ls, "\n")
	natsort.Sort(output_ls_slices)

	var final_output = output_ls_slices[len(output_ls_slices)-2][8:]
	return final_output
}
