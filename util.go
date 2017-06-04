package main

import (
	"os"
	"strings"

	"strconv"

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

func stringToKeys(input string) []int64 {
	var output []int64
	splitString := strings.Split(input, ",")
	for _, s := range splitString {
		s, _ := strconv.ParseInt(s, 10, 64)
		output = append(output, s)
	}
	return output
}
