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
	}

}

func CommandExit(args []string) error {
	fmt.Println("Closing the CEFI CLI... Goodbye!")
	return e.ErrExit
}
