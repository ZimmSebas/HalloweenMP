package user

import "fmt"

type User struct {
	Username    string
	Password    string
	Permissions PermStrategy
	Description string
}

type UserSystem struct {
	users map[string]User
}

func NewUserSystem() *UserSystem {
	return &UserSystem{
		users: map[string]User{
			"invitado":    *Guest(),
			"ines.vatela": *InesVatela(),
			"medium":      {Password: "", Permissions: &StaticPerm{}, Description: ""},
			"secretario":  {Password: "", Permissions: &StaticPerm{}, Description: ""},
			"cto":         {Password: "", Permissions: &StaticPerm{}, Description: ""},
			"jefeseg":     {Password: "", Permissions: &StaticPerm{}, Description: ""},
			"zordon":      {Password: "", Permissions: &StaticPerm{}, Description: ""},
			"gandalf":     {Password: "aguantesam", Permissions: &StaticPerm{}, Description: "Usuario administrador. Solo disponible para arreglos en el sistema."},
		},
	}
}

func (us *UserSystem) LoginAccess(username string, password string) (User, bool) {

	user, ok := us.users[username]

	if ok {
		return user, user.Password == password
	}

	guest_user := Guest()
	return *guest_user, false
}

func (us *UserSystem) ListUsers() (string, bool) {
	fmt.Println("Lista de usuarios:")
	fmt.Println("")

	for user := range us.users {
		fmt.Println("\t", user)
	}

	fmt.Println("")

	return "", true
}

func (us *UserSystem) GetDescription(username string) (string, bool) {

	user, ok := us.users[username]

	if ok {
		return user.Description, true
	}

	return "Usuario no encontrado", false
}

func Guest() *User {
	return &User{
		Username: "invitado",
		Password: "",
		Permissions: &StaticPerm{
			perms: PermissionSet{
				AllowedCommands: map[string]bool{
					"help":  true,
					"login": true,
				},
				AllowedFiles: map[string]bool{},
			},
		},
		Description: "Usuario invitado. \nSin permisos ni archivos.\n",
	}
}

func InesVatela() *User {
	return &User{
		Username: "ines.vatela",
		Password: "",
		Permissions: &StaticPerm{
			perms: PermissionSet{
				AllowedCommands: map[string]bool{
					"help":  true,
					"login": true,
				},
				AllowedFiles: map[string]bool{},
			},
		},
		Description: "Ines Vatela. PhD en Bioquímica aplicada a ecología. \n Coordinadora e investigadora del equipo de investigación ecológica. \n Nivel de seguridad: 2",
	}
}

type PermStrategy interface {
	CanAccessFile(filename string) bool
	CanSeeCommand(command_name string) bool
}

type PermissionSet struct {
	AllowedFiles    map[string]bool
	AllowedCommands map[string]bool
}

type StaticPerm struct {
	perms PermissionSet
}

func (p *StaticPerm) CanAccessFile(name string) bool {
	return p.perms.AllowedFiles[name]
}

func (p *StaticPerm) CanSeeCommand(name string) bool {
	return p.perms.AllowedFiles[name]
}
