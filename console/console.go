package console

import (
	"bufio"
	"fmt"
	"halloween/app"
	"halloween/ascii"
	"os"
	"strings"
)

type Console struct {
	commands map[string]Command
}

func NewConsole() *Console {
	return &Console{make(map[string]Command)}
}

func (c *Console) Register(cmd Command) {
	c.commands[cmd.Name()] = cmd
}

func (c *Console) RegisterDefaultCommands() {
	c.Register(&HelpCommand{})
	c.Register(&ClearCommand{})
	c.Register(&LoginCommand{})
	c.Register(&StatusCommand{})
	c.Register(&LogoutCommand{})
	c.Register(&ListPersonCommand{})
	c.Register(&CheckCommand{})
	c.Register(&ListFilesCommand{})
	c.Register(&OpenCommand{})
}

func (c *Console) Execute(raw_command string, ctx *app.Context) (string, bool) {
	command_array := strings.Fields(raw_command)

	var result string
	var ok_command bool

	cmdName := command_array[0]

	if cmd, ok := c.commands[cmdName]; ok {
		if len(command_array)-1 != cmd.Length() {
			result, ok_command = "Error de comando, se esperaban más argumentos. Ver *ayuda* para mas información.", false
		} else {
			result, ok_command = cmd.Execute(ctx, command_array[1:])
		}
	} else {
		result, ok_command = "Comando no reconocido", false
	}

	return result, ok_command

}

func (c *Console) Init(ctx *app.Context) {
	c.Execute("clear", ctx)

	fmt.Println("Modo de recuperación activo.")
	fmt.Println("")
	ascii.Oracle()
	fmt.Println("")
	fmt.Println("Puede escribir *ayuda* para ver comandos posibles, o escribir un comando para ejecutarlo")

	c.PromptLoop(ctx)

}

func (c *Console) PromptLoop(ctx *app.Context) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("> ")

	raw_input, _ := reader.ReadString('\n')

	result, ok_command := c.Execute(raw_input, ctx)

	if !ok_command {
		fmt.Println(result)
	}

	c.PromptLoop(ctx)

}
