package console

import (
	"fmt"
	"halloween/app"
	"halloween/ascii"
	"halloween/user"
	"os"
	"os/exec"
	"strconv"
)

type Command interface {
	Execute(ctx *app.Context, args []string) (string, bool)
	Name() string
	Length() int
}

type ClearCommand struct{}

func (c *ClearCommand) Name() string { return "clear" }
func (c *ClearCommand) Length() int  { return 0 }

func (c *ClearCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return "", true
}

type HelpCommand struct{}

func (h *HelpCommand) Name() string { return "ayuda" }
func (h *HelpCommand) Length() int  { return 0 }

func (h *HelpCommand) Execute(ctx *app.Context, args []string) (string, bool) {

	fmt.Println("login [usuario] [contraseña] - para iniciar sesión")
	fmt.Println("clear - para limpiar la consola")
	fmt.Println("estado - para ver el estado del sistema")
	fmt.Println("logout - para salir de la sesión")
	fmt.Println("listar - para revisar la lista de personas en el sistema")
	fmt.Println("consulta [persona] - para revisar información sobre una persona")
	fmt.Println("archivos - para ver lista de archivos")
	fmt.Println("abrir [archivo] - para abrir un archivo")

	return "", true
}

type StatusCommand struct{}

func (s *StatusCommand) Name() string { return "estado" }
func (s *StatusCommand) Length() int  { return 0 }

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
func (l *LoginCommand) Length() int  { return 2 }

func (l *LoginCommand) Execute(ctx *app.Context, args []string) (string, bool) {

	usersys := ctx.UserSystem
	user, pass := args[0], args[1]
	new_user, ok := usersys.LoginAccess(user, pass)

	if ok {
		fmt.Println("Acceso autorizado con usuario ", new_user.Username)
		ctx.User = &new_user
		return new_user.Username, true
	} else {
		ascii.You_shall_not_pass()
		return "", false
	}
}

type LogoutCommand struct{}

func (lo *LogoutCommand) Name() string { return "logout" }
func (lo *LogoutCommand) Length() int  { return 0 }

func (lo *LogoutCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	ctx.User = user.Guest()
	return "", true
}

type ListPersonCommand struct{}

func (lp *ListPersonCommand) Name() string { return "listar" }
func (lp *ListPersonCommand) Length() int  { return 0 }

func (lp *ListPersonCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	users := ctx.UserSystem.Users

	fmt.Println("Lista de usuarios:")
	for user := range users {
		fmt.Println(user)
	}
	return "", true
}

type CheckCommand struct{}

func (cc *CheckCommand) Name() string { return "consulta" }
func (cc *CheckCommand) Length() int  { return 1 }

func (cc *CheckCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	user := args[0]
	result, ok := ctx.UserSystem.GetDescription(user)
	return result, ok
}

type ListFilesCommand struct{}

func (f *ListFilesCommand) Name() string { return "archivos" }
func (f *ListFilesCommand) Length() int  { return 0 }

func (f *ListFilesCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	return "", true
}

type OpenCommand struct{}

func (o *OpenCommand) Name() string { return "abrir" }
func (o *OpenCommand) Length() int  { return 1 }

func (o *OpenCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	return "", true
}
