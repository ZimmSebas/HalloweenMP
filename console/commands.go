package console

import (
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
