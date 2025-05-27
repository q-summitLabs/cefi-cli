package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	cmd "github.com/q-summitLabs/cefi-cli/commands"
	e "github.com/q-summitLabs/cefi-cli/internal/errors"
)

const COMMAND_KEY string = "CEFI > "

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome to Central Finance - CLI!\n")

	for {

		fmt.Printf("%s", COMMAND_KEY)

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprint(os.Stderr, "shit broken")
			}
			fmt.Println("\nGoodbye")
			os.Exit(0)
		}

		fields := strings.Fields(scanner.Text())

		if len(fields) == 0 {
			fmt.Println("Please provide a command.")
			continue
		}

		cmdName, args := fields[0], fields[1:]

		val, ok := cmd.CommandTable[cmdName]

		if !ok {
			fmt.Println("Unknown command:", cmdName)
			continue
		}

		if err := val.Callback(args); err != nil {
			if errors.Is(err, e.ErrExit) {
				os.Exit(0)
			}
			fmt.Println("Issue with callback:", err)
			os.Exit(1)
		}
	}
}
