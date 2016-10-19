package main

import (
	"os/exec"
	"os"
	"log"
	"fmt"
)

var cmdChain = []*exec.Cmd{
	exec.Command("bin/synonyms"),
	exec.Command("bin/sprinkle"),
	exec.Command("bin/coolify"),
	exec.Command("bin/domainify"),
	exec.Command("bin/available"),
}

func main() {
	cmdChain[0].Stdin = os.Stdin
	cmdChain[len(cmdChain) - 1].Stdout = os.Stdout

	for i := 0; i < len(cmdChain) - 1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Fatalln(err)
		}
		nextCmd.Stdin = stdout
	}

	for _, cmd := range cmdChain {
		if err := cmd.Start(); err != nil {
			log.Fatalln("Could not start, " + err.Error())
		} else {
			defer cmd.Process.Kill()
		}
	}

	for index, cmd := range cmdChain {
		if err := cmd.Wait(); err != nil {
			log.Fatalln("Could not wait, index: " + fmt.Sprint(index) + ", " + err.Error())
		}
	}
}
