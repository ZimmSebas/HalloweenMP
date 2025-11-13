package user

type User struct {
	Username    string
	Password    string
	Permissions PermStrategy
	Description string
}

type UserSystem struct {
	Users map[string]User
}

func NewUserSystem() *UserSystem {
	return &UserSystem{
		Users: map[string]User{
			"guest":      {Username: "guest", Password: "", Permissions: &GuestPerm{}, Description: ""},
			"cientifica": {Username: "ines.vatela", Password: "", Permissions: &GuestPerm{}, Description: ""},
			"medium":     {Username: "medium", Password: "", Permissions: &GuestPerm{}, Description: ""},
			"secretario": {Username: "secretario", Password: "", Permissions: &GuestPerm{}, Description: ""},
			"cto":        {Username: "cto", Password: "", Permissions: &GuestPerm{}, Description: ""},
			"jefeseg":    {Username: "jefeseg", Password: "", Permissions: &GuestPerm{}, Description: ""},
			"zordon":     {Username: "manuel.zordon", Password: "", Permissions: &GuestPerm{}, Description: ""},
			"gandalf":    {Username: "gandalf", Password: "aguantesam", Permissions: &AdminPerm{}, Description: ""},
		},
	}
}

func (us *UserSystem) LoginAccess(username string, password string) (User, bool) {

	user, ok := us.Users[username]

	if ok {
		return user, user.Password == password
	}

	guest_user := Guest()
	return *guest_user, false
}

func (us *UserSystem) GetDescription(username string) (string, bool) {

	user, ok := us.Users[username]

	if ok {
		return user.Description, true
	}

	return "Usuario no encontrado", false
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
