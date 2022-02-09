package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/facette/natsort"
)

func main() {
	//Set terminal output colors
	const colorReset = "\033[0m"
	const colorRed = "\033[31m"
	const colorGreen = "\033[32m"

	//Get console outputs
	var running_kernel = cmdUname()
	var latest_kernel = cmdLsBoot()

	//Compare console outputs
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

	//Remove empty line (break) after running uname -r
	var output = strings.ReplaceAll(string(out), "\n", "")

	//CentOS7 sort fix
	var final_output = strings.ReplaceAll(output, ".el7.x86_64", "")

	return final_output
}

func cmdLsBoot() string {
	var cmd = "ls -1 /boot/ | grep \"vmlinuz\\|vmlinux\" | grep -v \"vmlinuz.old\\|vmlinux.old\""
	var out, err = exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	var output = strings.Split(string(out), "\n")

	var output_list = []string{}
	for _, _string := range output {
		//CentOS7 sort fix
		_string = strings.ReplaceAll(_string, ".el7.x86_64", "")

		if _string != "" {
			output_list = append(output_list, _string)
		}
	}
	natsort.Sort(output_list)

	var final_output = output_list[len(output_list)-1][8:]
	return final_output
}
