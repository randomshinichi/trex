package main

import (
	"github.com/olekukonko/tablewriter"
	"github.com/toorop/go-bittrex"
	"os"
	"strconv"
)

func fstring(f float64) string {
	return strconv.FormatFloat(f, 'f', 8, 64)
}

func main() {
	bittrex_api_key := os.Getenv("BITTREX_API_KEY")
	bittrex_api_secret := os.Getenv("BITTREX_API_SECRET")

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
}
