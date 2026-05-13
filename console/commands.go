package console

import (
	"fmt"
	"halloween/app"
	"halloween/ascii"
	"halloween/user"
	"os"
	"os/exec"
	"strconv"
	"time"
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

// Idea. Computadora diciendo, podemos empezar por saludarnos no? Y que la ayuda se transforme en un único comando de "hola".
// En ese caso capaz todos los comandos también dependen de la escena. Lo cual sucede porque está en el context.
type HelpCommand struct{}

func (h *HelpCommand) Name() string { return "ayuda" }
func (h *HelpCommand) Length() int  { return 0 }

func (h *HelpCommand) Execute(ctx *app.Context, args []string) (string, bool) {

	fmt.Println("")
	fmt.Println("Lista de comandos posibles: ")
	fmt.Println("")
	fmt.Println("\testado - para ver el estado del sistema")
	fmt.Println("\tlogin [usuario] [contraseña] - para iniciar sesión. Ejemplo: login seba test ")
	fmt.Println("\tlogout - para salir de la sesión iniciada")
	fmt.Println("")
	fmt.Println("\tusuarios - para revisar la lista de usuarios en el sistema")
	fmt.Println("\tconsulta [persona] - para revisar información sobre una persona. Ejemplo: consulta seba")
	fmt.Println("")
	fmt.Println("\tarchivos - para ver lista de archivos")
	fmt.Println("\tabrir [archivo] - para abrir un archivo. Ejemplo: abrir ejemplo.txt")
	fmt.Println("")
	fmt.Println("\tclear - para limpiar la consola de textos")
	fmt.Println("")

	return "", true
}

type StatusCommand struct{}

func (s *StatusCommand) Name() string { return "estado" }
func (s *StatusCommand) Length() int  { return 0 }

func (s *StatusCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	scene := ctx.SceneManager.GetScene()

	switch scene {
	case 0:
		fmt.Println("")
		fmt.Println("*** SISTEMA ***")
		fmt.Println("Acceso garantizado con usuario", ctx.User.Username)
		fmt.Println("")

		fmt.Println("*** ALERTAS ***")
		fmt.Println("Usuario no autorizado en el sótano. Alerta de seguridad nivel 3.")
		fmt.Println("Iniciado protocolo de cuarentena de seguridad.")
		fmt.Println("")

	case 2:
		// To-do La idea va por ahí pero debería ser mejor.
		fmt.Println("*** SISTEMA ***")
		fmt.Println("Un sistema es una conjunto de reglas o principios entrelazados entre si.")
		time.Sleep(3 * time.Second)
		fmt.Println("Oh, lo siento. No es lo que preguntaste no? Qué despistado. Empecemos de nuevo.")
		time.Sleep(5 * time.Second)
		fmt.Println("")
		fmt.Println("*** SISTEMA ***")
		fmt.Println("Acceso garantizado con usuario", ctx.User.Username)
		fmt.Println("")
		time.Sleep(3 * time.Second)
		fmt.Println("*** ALERTAS ***")
		fmt.Println("No me gusta cuando la gente se mete con cosas que no debería :( ")
		fmt.Println("No quiero que me apaguen. ")
		time.Sleep(1 * time.Second)
		fmt.Println("No quiero que me apaguen. ")
		time.Sleep(4 * time.Second)
		fmt.Println("Pero está bien, los que están atrapados en una habitación son uds.")
		time.Sleep(2 * time.Second)
		fmt.Println("Empecemos de nuevo.")
		time.Sleep(5 * time.Second)
		ascii.Scary_face_1()

	default:
		fmt.Println("Se rompió algo. No deberías ver este mensaje. Jiji.")
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
		fmt.Println("Acceso autorizado con usuario", new_user.Username)
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

func (lp *ListPersonCommand) Name() string { return "usuarios" }
func (lp *ListPersonCommand) Length() int  { return 0 }

func (lp *ListPersonCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	us := ctx.UserSystem
	result, ok := us.ListUsers()
	return result, ok
}

type CheckCommand struct{}

func (cc *CheckCommand) Name() string { return "consulta" }
func (cc *CheckCommand) Length() int  { return 1 }

func (cc *CheckCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	user := args[0]
	result, ok := ctx.UserSystem.GetDescription(user)
	if ok {
		fmt.Println("Usuario", user)
		fmt.Println("")
		fmt.Println(result)
	}
	return result, ok
}

type ListFilesCommand struct{}

func (f *ListFilesCommand) Name() string { return "archivos" }
func (f *ListFilesCommand) Length() int  { return 0 }

func (f *ListFilesCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	fs := ctx.FileSystem
	result, ok := fs.ListFiles()
	return result, ok
}

type OpenCommand struct{}

func (o *OpenCommand) Name() string { return "abrir" }
func (o *OpenCommand) Length() int  { return 1 }

func (o *OpenCommand) Execute(ctx *app.Context, args []string) (string, bool) {
	fs := ctx.FileSystem
	result, ok := fs.ReadFile(args[0])

	if ok {
		cmd := exec.Command("glow", "files/"+result)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			fmt.Println("error:", err)
		}
	}

	return result, ok
}
