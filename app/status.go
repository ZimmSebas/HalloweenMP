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
	Filesys *files.FileSystem
	Scene   *scene.SceneManager
	User    *user.User
}
