package main

import (
	"strconv"

	"github.com/vscale/go-vscale"
)

func listKeys() {
	client := vscale_api_go.NewClient(token)
	keys, _, _ := client.SSHKey.List()
	var data [][]string

	for _, k := range *keys {
		data = append(data, []string{strconv.Itoa(int(k.ID)), k.Name})
	}
	header := []string{"id", "name"}
	printTable(header, data)
}
