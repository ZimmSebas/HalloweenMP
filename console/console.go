package console

import (
	"bufio"
	"fmt"
	"halloween/app"
	"os"
	"strings"
)

type Console struct {
	commands map[string]Command
}

func (c *Console) Execute(raw_command string, ctx *app.Context) (string, bool) {
	command_array := strings.Fields(raw_command)

	var result string
	var ok_command bool

	cmdName := command_array[0]
	if cmd, ok := c.commands[cmdName]; ok {
		result, ok_command = cmd.Execute(ctx, command_array[1:])
	} else {
		result, ok_command = "Comando no reconocido", false
	}

	return result, ok_command

}

func (c *Console) Prompt(ctx *app.Context) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> ")

	raw_input, _ := reader.ReadString('\n')

	c.Execute(raw_input, ctx)
}
