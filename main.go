package main

import (
	"halloween/app"
	"halloween/console"
	"halloween/user"
)

func main() {

	u := user.Guest()

	ctx := app.NewContext(u)

	c := console.NewConsole()
	c.RegisterDefaultCommands()

	ctx.FileSystem.ListFiles()
}
