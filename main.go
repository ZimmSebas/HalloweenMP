package main

import (
	"halloween/app"
	"halloween/console"
	"halloween/user"
)

//func start(ctx *app.Context) {
//	ctx.SceneManager.StartScene()
//}

func main() {

	u := user.Guest()

	ctx := app.NewContext(u)

	c := console.NewConsole()
	c.RegisterDefaultCommands()

	ctx.FileSystem.ListFiles()

	//start(ctx)
	c.Init(ctx)

}
