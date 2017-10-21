package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/toorop/go-bittrex"
	"os"
	"strconv"
	"time"
)

func fstring(f float64) string {
	return strconv.FormatFloat(f, 'f', 8, 64)
}

func main() {
	bittrex_api_key := os.Getenv("BITTREX_API_KEY")
	bittrex_api_secret := os.Getenv("BITTREX_API_SECRET")

	bittrex := bittrex.New(bittrex_api_key, bittrex_api_secret)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Timestamp", "OrderType", "Limit", "Quantity", "Price", "PricePerUnit"})

	args := os.Args[1:]
	market := fmt.Sprintf("%v-%v", args[0], args[1])
	order_history, _ := bittrex.GetOrderHistory(market)

	for _, order := range order_history {
		ts := order.TimeStamp.Format(time.UnixDate)
		s := []string{ts, order.OrderType, fstring(order.Limit), fstring(order.Quantity), fstring(order.Price), fstring(order.PricePerUnit)}
		table.Append(s)
	}
	table.Render()
}
