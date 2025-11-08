package user

type PermStrategy interface {
	CanAccessFile(filename string) bool
}

type GuestPerm struct{}

func (g *GuestPerm) CanAccessFile(name string) bool {

	// See how in hell to make work permissions.

	return false
}

type User struct {
	Name        string
	Permissions PermStrategy
}
