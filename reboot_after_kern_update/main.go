package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"strings"

	"github.com/facette/natsort"
)

type json_response_bad struct {
	Status         int
	Running_kernel string
	Latest_kernel  string
}

func jsonResposeBad(status int, running_kernel, latest_kernel string) string {
	var json_response = &json_response_bad{
		Status:         status,
		Running_kernel: running_kernel,
		Latest_kernel:  latest_kernel,
	}

	var final_json_response, _ = json.Marshal(json_response)
	return string(final_json_response)
}

type json_response_good struct {
	Status         int
	Running_kernel string
}

func jsonResposeGood(status int, running_kernel string) string {
	var json_response = &json_response_good{
		Status:         status,
		Running_kernel: running_kernel,
	}

	var final_json_response, _ = json.Marshal(json_response)
	return string(final_json_response)
}

func main() {
	//Set terminal output colors
	const colorReset = "\033[0m"
	const colorRed = "\033[31m"
	const colorGreen = "\033[32m"

	//Get command line flags
	var jsonOutputFlag = flag.Bool("json", false, "Use JSON output")
	flag.Parse()

	//Run 2 functions to get console outputs for "uname -r" and "ls /boot"
	var running_kernel = cmdUname()
	var latest_kernel = cmdLsBoot()

	//Init status variable for later
	var status int

	//Compare console outputs
	if running_kernel == latest_kernel {
		status = 0
		if *jsonOutputFlag {
			var jsonRespose = jsonResposeGood(status, running_kernel)
			fmt.Println(jsonRespose)
		} else {
			fmt.Println(colorGreen + " 🟢 All good. You are running the latest kernel version." + colorReset)
		}
	} else {
		status = 1
		if *jsonOutputFlag {
			var jsonRespose = jsonResposeBad(status, running_kernel, latest_kernel)
			fmt.Println(jsonRespose)
		} else {
			fmt.Println(colorRed + " 🔴 Please reboot to apply the kernel update!" + colorReset)
			fmt.Println()
			fmt.Println("You are running:      " + colorRed + running_kernel + colorReset)
			fmt.Println("The latest installed: " + colorGreen + latest_kernel + colorReset)
		}
	}
}

func cmdUname() string {
	//Get command line output
	var cmd = "uname -r"
	var out, err = exec.Command("bash", "-c", cmd).Output()

	//Handle errors
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
	//Get command line output
	var cmd = "ls -1 /boot/ | grep \"vmlinuz\\|vmlinux\" | grep -v \"vmlinuz.old\\|vmlinux.old\""
	var out, err = exec.Command("bash", "-c", cmd).Output()

	//Handle errors
	if err != nil {
		fmt.Printf("%s", err)
	}

	//Parse the output
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
