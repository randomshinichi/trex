package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

// var (
//     balanceCommand = cli.Command{
//         Name: "ba",
//         Usage: "List Balances",
//         ArgsUsage: "",
//         Category: "NORMAL COMMANDS",
//         Description: `trex ba will list your balances`,
//     }
// )
func main() {
	app := cli.NewApp()
	app.Name = "trex"
	app.Usage = "easy interface with bittrex"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Boom! I say!")
		return nil
	}

	app.Run(os.Args)
}
