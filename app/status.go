package app

import (
	"halloween/files"
	"halloween/scene"
	"halloween/user"
)

type Mediator interface {
	GetFileSystem() *files.FileSystem
	GetSceneManager() *scene.SceneManager
	GetUserSystem() *user.UserSystem
	GetUser() *user.User
}

type Context struct {
	FileSystem   *files.FileSystem
	SceneManager *scene.SceneManager
	UserSystem   *user.UserSystem
	User         *user.User
}

func NewContext(u *user.User) *Context {
	return &Context{
		FileSystem:   files.NewFileSystem(),
		SceneManager: scene.NewSceneManager(),
		UserSystem:   user.NewUserSystem(),
		User:         u,
	}
}
