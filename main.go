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
	table.Render()
}
