package user

type User struct {
	Username    string
	Password    string
	Permissions PermStrategy
}

type UserSystem struct {
	users map[string]User
}

func NewUserSystem() *UserSystem {
	return &UserSystem{
		users: map[string]User{
			"guest":      {Username: "guest", Password: "", Permissions: &GuestPerm{}},
			"cientifica": {Username: "ines.vatela", Password: "", Permissions: &GuestPerm{}},
			"medium":     {Username: "medium", Password: "", Permissions: &GuestPerm{}},
			"secretario": {Username: "secretario", Password: "", Permissions: &GuestPerm{}},
			"cto":        {Username: "cto", Password: "", Permissions: &GuestPerm{}},
			"jefeseg":    {Username: "jefeseg", Password: "", Permissions: &GuestPerm{}},
			"zordon":     {Username: "manuel.zordon", Password: "", Permissions: &GuestPerm{}},
			"gandalf":    {Username: "gandalf", Password: "aguantesam", Permissions: &AdminPerm{}},
		},
	}
}

func Guest() *User {
	return &User{
		Username:    "guest",
		Password:    "",
		Permissions: &GuestPerm{},
	}
}

type PermStrategy interface {
	CanAccessFile(filename string) bool
	CanSeeCommand(command_name string) bool
}

type GuestPerm struct{}

func (g *GuestPerm) CanAccessFile(name string) bool {

	// See how in hell to make work permissions.

	return false
}

func (g *GuestPerm) CanSeeCommand(name string) bool {

	// See how in hell to make work permissions.

	return false
}

type AdminPerm struct{}

func (a *AdminPerm) CanAccessFile(name string) bool {
	return true
}

func (a *AdminPerm) CanSeeCommand(name string) bool {
	return true
}
