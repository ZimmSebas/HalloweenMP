package app

import (
	"halloween/files"
	"halloween/scene"
	"halloween/user"
)

type Mediator interface {
	GetFileSystem() *files.FileSystem
	GetSceneManager() *scene.SceneManager
	GetUser() *user.User
}

type Context struct {
	FileSystem   *files.FileSystem
	SceneManager *scene.SceneManager
	User         *user.User
}

func NewContext(u *user.User) *Context {
	return &Context{
		FileSystem:   files.NewFileSystem(),
		SceneManager: scene.NewSceneManager(),
		User:         u,
	}
}
