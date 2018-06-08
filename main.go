package main

import (
	"github.com/urfave/cli"
	"neo-go-compiler/cli/smartcontract"
	"os"
)

func main() {
	ctl := cli.NewApp()
	ctl.Name = "neo-go-compiler"
	ctl.Usage = "golang compiler for neo smart contract"

	ctl.Commands = []cli.Command{
		smartcontract.NewCommand(),
	}

	ctl.Run(os.Args)
}
