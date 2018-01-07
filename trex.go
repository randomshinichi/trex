package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var (
	balanceCommand = cli.Command{
		Name:    "balance",
		Aliases: []string{"ba"},
		Usage:   "List Balances",
		Action: func(c *cli.Context) error {
			fmt.Println("balance")
			return nil
		},
	}
	ordersCommand = cli.Command{
		Name:  "orders",
		Usage: "List open orders",
		Action: func(c *cli.Context) error {
			fmt.Println("orders")
			return nil
		},
	}
	ordershistCommand = cli.Command{
		Name:  "ordershist",
		Usage: "List historical orders",
		Action: func(c *cli.Context) error {
			fmt.Println("orders_hist")
			return nil
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "trex"
	app.Usage = "easy interface with bittrex"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "bittrex token",
			EnvVar: "BITTREX_API_KEY",
		},
		cli.StringFlag{
			Name:   "secret",
			Usage:  "bittrex secret",
			EnvVar: "BITTREX_API_SECRET",
		},
	}
	app.Commands = []cli.Command{balanceCommand, ordersCommand, ordershistCommand}
	app.Run(os.Args)
}
