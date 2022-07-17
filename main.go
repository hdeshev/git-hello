package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Fruit", "Amount"})
	table.Append([]string{"apples", "10"})
	table.Append([]string{"pears", "30"})
	table.Append([]string{"raspberries", "3"})
	table.Append([]string{"plums", "5"})
	table.Append([]string{"cherries", "2"})
	table.Append([]string{"bananas", "8"})
	table.Render()
}
