package commands

import (
	"fmt"

	e "github.com/q-summitLabs/cefi-cli/internal/errors"
	"github.com/q-summitLabs/cefi-cli/models"
)

var CommandTable map[string]models.Command

func init() {

	CommandTable = map[string]models.Command{
		"exit": {
			Name:        "exit",
			Description: "Exits the CLI",
			Callback:    func(args []string) error { return CommandExit(args) },
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func(args []string) error { return CommandHelp(args) },
		},
	}

}

func CommandExit(args []string) error {
	if len(args) > 0 {
		return e.ErrTooManyArgs
	}
	fmt.Println("Closing the CEFI CLI... Goodbye!")
	return e.ErrExit
}

func CommandHelp(args []string) error {
	if len(args) > 0 {
		return e.ErrTooManyArgs
	}

	fmt.Println()
	fmt.Println("Usage:")

	for cmd, val := range CommandTable {
		fmt.Printf("  %s: %v\n", cmd, val.Description)
	}

	fmt.Println()

	return nil
}
