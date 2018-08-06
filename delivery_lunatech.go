package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func logMessage(message string) string {
	dateNow := time.Now()
	dtFormated := dateNow.Format("2006-01-02 15:04:05.000")
	msg := "[ " + dtFormated + " ] " + message
	return msg
}

func checkArgs() (string, string, string) {
	if argsLen := len(os.Args); argsLen > 1 {
		app := os.Args[1]
		version := os.Args[2]
		script := "redeploy.sh"
		fmt.Println(logMessage("Redeploying the " + app + " in " + version + " version."))
		return script, app, version
	}
	script := "deploy.sh"
	fmt.Println(logMessage("Executing the deploy of the entire stack."))
	return script, "", ""
}

func getOs() string {
	if osVersion := runtime.GOOS; osVersion == "linux" {
		return osVersion
	} else if osVersion == "windows" {
		return osVersion
	}
	return "Unknown"
}

func operation(osVersion string, script string, app string, version string) {

	cmd := exec.Command(osVersion+"/"+script, app, version)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf(logMessage("%s %s\n"), "Error creating StdoutPipe for Cmd", err)
		return
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Printf(logMessage("%s %s\n"), "Error starting Cmd", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf(logMessage("%s %s\n"), "Error waiting for Cmd", err)
		return
	}
}

func main() {
	script, app, version := checkArgs()
	os := getOs()
	operation(os, script, app, version)

}
