package main

import (
	"github.com/urfave/cli"
	"github.com/ontio/neo-go-compiler/cli/smartcontract"
	"os"
)

func main() {
	ctl := cli.NewApp()
	ctl.Name = "neo-go-compiler"
	ctl.Usage = "golang compiler for neo smart contract"
	ctl.Version = "0.1.0"

	ctl.Commands = []cli.Command{
		smartcontract.NewCommand(),
	}

	ctl.Run(os.Args)
}
