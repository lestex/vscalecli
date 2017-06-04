package main

import (
	"strconv"

	"github.com/vscale/go-vscale"
)

func listRplans() {
	client := vscale_api_go.NewClient(token)
	rplans, _, _ := client.Configuration.ListRplans()
	var data [][]string

	for _, rp := range *rplans {
		data = append(data, []string{rp.ID,
			strconv.Itoa(int(rp.CPUs)),
			strconv.Itoa(int(rp.Disk) / 1024),
			strconv.Itoa(int(rp.Memory)),
			strconv.Itoa(int(rp.Network))})
	}
	header := []string{"id",
		"cpus",
		"disk",
		"Memory",
		"Network"}
	printTable(header, data)
}
