package main

import (
	"errors"
	"os"

	"github.com/olekukonko/tablewriter"
)

func getToken() (string, error) {

	token := os.Getenv("VSCALE")
	if token == "" {
		return token, errors.New("Token is empty")
	}

	return token, nil
}

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
