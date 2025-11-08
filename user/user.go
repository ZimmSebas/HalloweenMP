package user

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

type User struct {
	Name        string
	Permissions PermStrategy
}

func Guest() *User {
	return &User{
		Name:        "guest",
		Permissions: &GuestPerm{},
	}
}
