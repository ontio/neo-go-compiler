package smartcontract

import (
	"github.com/urfave/cli"
	"neo-go-compiler/compiler"
)

const (
	errNoInput = "Input file is mandatory and should be passed using -i flag."
)

// NewCommand returns a new contract command.
func NewCommand() cli.Command {
	return cli.Command{
		Name:  "contract",
		Usage: "compile smart contracts",
		Subcommands: []cli.Command{
			{
				Name:   "compile",
				Usage:  "compile a smart contract to a .avm file",
				Action: contractCompile,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "in, i",
						Usage: "Input file for the smart contract to be compiled",
					},
					cli.StringFlag{
						Name:  "out, o",
						Usage: "Output of the compiled contract",
					},
					cli.BoolFlag{
						Name:  "debug, d",
						Usage: "Debug mode will print out additional information after a compiling",
					},
					cli.BoolFlag{
						Name:  "abi, a",
						Usage: "generate abi file",
					},
				},
			},
			{
				Name:   "opdump",
				Usage:  "dump the opcode of a .go file",
				Action: contractDumpOpcode,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "in, i",
						Usage: "Input file for the smart contract",
					},
				},
			},
		},
	}
}

func contractCompile(ctx *cli.Context) error {
	src := ctx.String("in")
	if len(src) == 0 {
		return cli.NewExitError(errNoInput, 1)
	}

	o := &compiler.Options{
		Outfile: ctx.String("out"),
		Debug:   ctx.Bool("debug"),
		AbiGen:  ctx.Bool("abi"),
	}

	if err := compiler.CompileAndSave(src, o); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

//
//func testInvoke(ctx *cli.Context) error {
//	src := ctx.String("in")
//	if len(src) == 0 {
//		return cli.NewExitError(errNoInput, 1)
//	}
//
//	b, err := ioutil.ReadFile(src)
//	if err != nil {
//		return cli.NewExitError(err, 1)
//	}
//
//	// For now we will hardcode the endpoint.
//	// On the long term the internal VM will run the script.
//	// TODO: remove RPC dependency, hardcoded node.
//	endpoint := "http://seed5.bridgeprotocol.io:10332"
//	opts := rpc.ClientOptions{}
//	client, err := rpc.NewClient(context.TODO(), endpoint, opts)
//	if err != nil {
//		return cli.NewExitError(err, 1)
//	}
//
//	scriptHex := hex.EncodeToString(b)
//	resp, err := client.InvokeScript(scriptHex)
//	if err != nil {
//		return cli.NewExitError(err, 1)
//	}
//
//	b, err = json.MarshalIndent(resp.Result, "", "  ")
//	if err != nil {
//		return cli.NewExitError(err, 1)
//	}
//
//	fmt.Println(string(b))
//
//	return nil
//}

func contractDumpOpcode(ctx *cli.Context) error {
	src := ctx.String("in")
	if len(src) == 0 {
		return cli.NewExitError(errNoInput, 1)
	}
	if err := compiler.DumpOpcode(src); err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}
