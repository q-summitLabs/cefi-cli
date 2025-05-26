package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		}
	}
}
