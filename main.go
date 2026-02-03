package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")

		raw_string, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			os.Exit(1)
		}

		trimmed := strings.TrimSpace(raw_string)
		command, args := parse_command(trimmed)

		switch command {

		case "":
			fmt.Println()

		case "type":
			if len(args) == 0 {
				fmt.Println()
			} else {
				if _, exists := get_method_bound_to_command(args[0]); exists {
					fmt.Println(args[0] + " is a shell builtin")
				} else {
					fmt.Println(args[0] + ": not found")
				}
			}

		default:
			comand_func, exists := get_method_bound_to_command(command)

			if exists {
				comand_func(args...)
			} else {
				fmt.Println(command + ": command not found")
			}

		}

	}
}

type commands map[string]func(args ...string)

var known_commands = commands{
	"exit": func(args ...string) { os.Exit(0) },

	"echo": func(args ...string) { fmt.Println(strings.Join(args, " ")) },

	"type": func(args ...string) { /* returns the type (done separately) */ },
}

func get_method_bound_to_command(command string) (func(args ...string), bool) {
	comand_func, exists := known_commands[command]
	return comand_func, exists
}

func parse_command(input string) (string, []string) {
	parts := strings.Split(input, " ")

	if len(parts) > 1 {
		return parts[0], parts[1:]
	}

	return parts[0], nil
}

