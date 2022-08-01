package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/facette/natsort"
)

type json_response_bad struct {
	Status         int    `json:"status"`
	Running_kernel string `json:"running_kernel"`
	Latest_kernel  string `json:"latest_kernel"`
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
	Status         int    `json:"status"`
	Running_kernel string `json:"running_kernel"`
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
			fmt.Println(colorGreen + " 🟢 " + colorReset + "All good. You are running the latest kernel version.")
		}
	} else {
		status = 1
		if *jsonOutputFlag {
			var jsonRespose = jsonResposeBad(status, running_kernel, latest_kernel)
			fmt.Println(jsonRespose)
		} else {
			fmt.Println(colorRed + " 🔴 " + colorReset + "Please reboot to apply the kernel update!")
			fmt.Println("       Currently active kernel:     " + running_kernel)
			fmt.Println("       The latest installed:        " + colorGreen + latest_kernel + colorReset)
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
	if err != nil {
		fmt.Printf("%s", err)
	}

	//Parse the output
	var output = strings.Split(string(out), "\n")
	var output_list []string
	var output_list_oem []string
	var output_list_normal []string

	// Detect if OEM kernel is in use
	var r, _ = regexp.Compile(".*oem.*")
	for _, i := range output {
		var reMatch = r.MatchString(i)
		if reMatch {
			output_list_oem = append(output_list_oem, i)
		} else {
			output_list_normal = append(output_list_normal, i)
		}
	}
	fmt.Println("OEM list: ", output_list_oem)
	fmt.Println("Normal list: ", output_list_normal)

	// Use EOM kernel if detected
	var unameVar = cmdUname()
	if r.MatchString(unameVar) {
		output_list = output_list_oem
	} else {
		output_list = output_list_normal
	}
	fmt.Println("Final list: ", output_list)

	for _, i := range output_list {
		//CentOS7 sort fix
		i = strings.ReplaceAll(i, ".el7.x86_64", "")

		if i != "" {
			output_list = append(output_list, i)
		}
	}
	natsort.Sort(output_list)
	fmt.Println("Final sorted list: ", output_list_normal)

	var final_output = output_list[len(output_list)-1]
	final_output = strings.ReplaceAll(final_output, "vmlinuz-", "")

	fmt.Println("Final single item: ", final_output)
	return final_output
}
