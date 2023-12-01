package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var comm string
	var workingDirectory string
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		comm, _ = reader.ReadString('\n')
		// Prepare command
		comm = strings.TrimSpace(comm)
		command := strings.Split(comm, " ")
		if comm == "stop" {
			return
		}

		workingDirectory, _ = os.Getwd()

		switch command[0] {
		case "cd":
			os.Chdir(command[1])
			workingDirectory = command[1]
			continue
		case "pwd":
			p, err := os.Getwd()
			if err != nil {
				fmt.Println(os.Stderr, err)
			}
			fmt.Println(p)
			continue
		case "echo":
			fmt.Println(command[1])
			break
		default:
			cmd := exec.Command(command[0], command[1:]...)
			cmd.Dir = workingDirectory
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Print(string(out))
		}

	}
}
