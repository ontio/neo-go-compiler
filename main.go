package main

import (
	"os"
	"github.com/urfave/cli"
	"neo-go-compiler/cli/smartcontract"
)


func main() {
	ctl := cli.NewApp()
	ctl.Name = "neo-go"
	ctl.Usage = "Official Go client for Neo"

	ctl.Commands = []cli.Command{
		smartcontract.NewCommand(),
	}

	ctl.Run(os.Args)
}
