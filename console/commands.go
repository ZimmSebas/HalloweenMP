package console

import (
	"fmt"
	"halloween/app"
	"halloween/ascii"
	"os"
	"os/exec"
	"strconv"
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
func (h *HelpCommand) Execute(ctx *app.Context, args []string) (string, bool) {

	fmt.Println("login [usuario] - para iniciar sesión")
	fmt.Println("estado - para ver el estado del sistema")
	fmt.Println("logout - para salir de la sesión")
	fmt.Println("consulta [nombre] - para revisar información sobre una persona")
	fmt.Println("archivos - para ver lista de archivos")
	fmt.Println("abrir [archivo] - para abrir un archivo")

	return "", true
}

type StatusCommand struct{}

func (s *StatusCommand) Name() string { return "estado" }
func (s *StatusCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	scene := ctx.SceneManager.GetScene()

	switch scene {
	case 0:
		fmt.Println("Usuario no autorizado en el sótano. Alerta de seguridad nivel 3.")
		fmt.Println("Iniciando protocolo de cuarentena de seguridad.")

	default:
		fmt.Println("Se rompió algo.")
	}
	return strconv.Itoa(scene), true
}

type LoginCommand struct{}

func (l *LoginCommand) Name() string { return "login" }
func (l *LoginCommand) Execute(ctx *app.Context, args []string) (string, bool) {

	usersys := ctx.UserSystem
	user, pass := args[1], args[2]
	new_user, ok := usersys.CanAccess(user, pass)

	if ok {
		fmt.Println("Acceso autorizado con usuario ", new_user.Username)
		ctx.User = &new_user
		return new_user.Username, true
	} else {
		ascii.You_shall_not_pass()
		return "", false
	}
}
