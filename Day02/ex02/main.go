package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args
	length := len(args)
	if length == 1 {
		fmt.Println("Error: no arguments present.")
		os.Exit(1)
	}
	command := args[1]
	var arguments []string
	if length == 2 {
		arguments = append(arguments, "")
	} else {
		arguments = append(arguments, args[2:]...)
	}
	read := bufio.NewScanner(os.Stdin)
	for read.Scan() {
		str := strings.Split(read.Text(), "\n ")
		for _, arg := range str {
			fullArgs := append(arguments, arg)
			cmd := exec.Command(command, fullArgs...)
			res, err := cmd.CombinedOutput()
			if err == nil {
				fmt.Println(string(res))
			}
		}
	}
}
