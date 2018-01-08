package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/toorop/go-bittrex"
	"gopkg.in/urfave/cli.v1"
	"os"
	"time"
	"github.com/shopspring/decimal"
)

func fstring(f decimal.Decimal) string {
	return f.String()
}

func marketFromArgs(c *cli.Context) string {
	var market string
	fmt.Printf("%#v\n", c)
	if c.NArg() > 0 {
		market = fmt.Sprintf("%v-%v", c.Args().Get(0), c.Args().Get(1))
	} else{
		market = "all"
	}
	return market
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
		Aliases: []string{"b"},
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
				if !balance.Balance.Equals(decimal.Zero) {
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
			table.SetHeader([]string{"OrderUUID","OrderType", "Exchange", "Limit", "Quantity", "Price", "PricePerUnit"})

			order_history, _ := bittrex.GetOpenOrders(marketFromArgs(c))

			for _, order := range order_history {
				s := []string{order.OrderUuid, order.OrderType, order.Exchange,fstring(order.Limit), fstring(order.Quantity), fstring(order.Price), fstring(order.PricePerUnit)}
				table.Append(s)
			}
			table.Render()
			return nil
		},
	}
	histCommand = cli.Command{
		Name:  "hist",
		Usage: "hist <BASE> <ALT>",
		Flags: []cli.Flag{tokenFlag, secretFlag},
		Action: func(c *cli.Context) error {
			bittrex_api_key := c.String("key")
			bittrex_api_secret := c.String("secret")

			bittrex := bittrex.New(bittrex_api_key, bittrex_api_secret)
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Timestamp", "Exchange", "OrderType", "Limit", "Quantity", "Price", "PricePerUnit"})

			order_history, _ := bittrex.GetOrderHistory(marketFromArgs(c))

			for _, order := range order_history {
				ts := order.TimeStamp.Format(time.UnixDate)
				s := []string{ts, order.Exchange, order.OrderType, fstring(order.Limit), fstring(order.Quantity), fstring(order.Price), fstring(order.PricePerUnit)}
				table.Append(s)
			}
			table.Render()
			return nil
		},
	}
	cancelCommand = cli.Command{
		Name:"cancel",
		Usage: "cancel OrderUUID",
		Flags: []cli.Flag{tokenFlag, secretFlag},
		Action: func(c *cli.Context) error {
			bittrex_api_key := c.String("key")
			bittrex_api_secret := c.String("secret")

			bittrex := bittrex.New(bittrex_api_key, bittrex_api_secret)
			err := bittrex.CancelOrder(c.Args().Get(0))
			if err != nil {
				fmt.Printf("%v\n", err)
				return err
			}
			return nil
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "trex"
	app.Usage = "easy interface with bittrex"
	app.Commands = []cli.Command{balanceCommand, ordersCommand, histCommand, cancelCommand}
	app.Run(os.Args)
}
