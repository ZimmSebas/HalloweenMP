package console

import (
	"fmt"
	"halloween/app"
	"os"
	"os/exec"
)

type Command interface {
	Execute(ctx *app.Context, args []string) (string, bool)
	Name() string
}

type ClearCommand struct{}

func (c *ClearCommand) Name() string { return "clear" }
func (c *ClearCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return "", true
}

type HelpCommand struct{}

func (h *HelpCommand) Name() string { return "help" }
func (c *HelpCommand) Execute(ctx *app.Context, args []string) (string, bool) {

	fmt.Println("login [usuario] - para iniciar sesión")
	fmt.Println("estado - para ver el estado del sistema")
	fmt.Println("logout - para salir de la sesión")
	fmt.Println("consulta [nombre] - para revisar información sobre una persona")
	fmt.Println("archivos - para ver lista de archivos")
	fmt.Println("abrir [archivo] - para abrir un archivo")

	return "", true
}
