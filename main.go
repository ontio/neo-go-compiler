package main

import (
	"github.com/urfave/cli"
	"neo-go-compiler/cli/smartcontract"
	"os"
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
