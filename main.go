package main

import (
	"bwastartup/app"
	"bwastartup/cli"
	"os"
)

func main() {

	c := cli.NewCli(os.Args)
	c.Run(app.Init())

}
