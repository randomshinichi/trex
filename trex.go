package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/toorop/go-bittrex"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"time"
)

func fstring(f float64) string {
	return strconv.FormatFloat(f, 'f', 8, 64)
}

var (
	tokenFlag = cli.StringFlag{
		Name:   "key",
		Usage:  "bittrex api key",
		EnvVar: "BITTREX_API_KEY",
	}
	secretFlag = cli.StringFlag{
		Name:   "secret",
		Usage:  "bittrex api secret",
		EnvVar: "BITTREX_API_SECRET",
	}

	balanceCommand = cli.Command{
		Name:    "balance",
		Aliases: []string{"ba"},
		Usage:   "List Balances",
		Flags:   []cli.Flag{tokenFlag, secretFlag},
		Action: func(c *cli.Context) error {
			bittrex_api_key := c.String("key")
			bittrex_api_secret := c.String("secret")
			bittrex := bittrex.New(bittrex_api_key, bittrex_api_secret)

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Coin", "Total", "Not In Orders", "Pending Deposit"})

			balances, _ := bittrex.GetBalances()
			for _, balance := range balances {
				s := []string{balance.Currency, fstring(balance.Balance), fstring(balance.Available), fstring(balance.Pending)}
				if balance.Balance != 0 {
					table.Append(s)
				}
			}
			table.Render()
			return nil
		},
	}
	ordersCommand = cli.Command{
		Name:  "orders",
		Usage: "List open orders for BASE ALT",
		Flags: []cli.Flag{tokenFlag, secretFlag},
		Action: func(c *cli.Context) error {
			bittrex_api_key := c.String("key")
			bittrex_api_secret := c.String("secret")

			bittrex := bittrex.New(bittrex_api_key, bittrex_api_secret)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"OrderType", "Limit", "Quantity", "Price", "PricePerUnit"})

			market := fmt.Sprintf("%v-%v", c.Args().Get(0), c.Args().Get(1))
			order_history, _ := bittrex.GetOpenOrders(market)

			for _, order := range order_history {
				s := []string{order.OrderType, fstring(order.Limit), fstring(order.Quantity), fstring(order.Price), fstring(order.PricePerUnit)}
				table.Append(s)
			}
			table.Render()
			return nil
		},
	}
	ordershistCommand = cli.Command{
		Name:  "ordershist",
		Usage: "List historical orders",
		Flags: []cli.Flag{tokenFlag, secretFlag},
		Action: func(c *cli.Context) error {
			bittrex_api_key := c.String("key")
			bittrex_api_secret := c.String("secret")

			bittrex := bittrex.New(bittrex_api_key, bittrex_api_secret)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Timestamp", "Exchange", "OrderType", "Limit", "Quantity", "Price", "PricePerUnit"})

			var market string
			if c.NArg() > 0 {
				market = fmt.Sprintf("%v-%v", c.Args().Get(0), c.Args().Get(1))
			} else {
				market = "all"
			}
			order_history, _ := bittrex.GetOrderHistory(market)

			for _, order := range order_history {
				ts := order.TimeStamp.Format(time.UnixDate)
				s := []string{ts, order.Exchange, order.OrderType, fstring(order.Limit), fstring(order.Quantity), fstring(order.Price), fstring(order.PricePerUnit)}
				table.Append(s)
			}
			table.Render()
			return nil
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "trex"
	app.Usage = "easy interface with bittrex"
	app.Commands = []cli.Command{balanceCommand, ordersCommand, ordershistCommand}
	app.Run(os.Args)
}
