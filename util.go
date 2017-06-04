package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func printTable(header []string, data [][]string) {
	if data == nil {
		return
	}
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	for _, d := range data {
		table.Append(d)
	}
	table.Render()
}
