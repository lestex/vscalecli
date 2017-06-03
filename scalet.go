package main

import (
	"strconv"

	"github.com/vscale/go-vscale"
)

// (makeFrom, rplan, name, password, location string, doStart bool,keys []int64, wait bool)
func scaletCreate() {

}

func scaletList() {
	client := vscale_api_go.NewClient(token)
	scalets, _, _ := client.Scalet.List()
	var data [][]string

	for _, s := range *scalets {
		data = append(data, []string{strconv.Itoa(int(s.CTID)), s.Hostname, s.MadeFrom, s.Rplan, s.Location, s.PublicAddresses.Address, s.Rplan})
	}
	header := []string{"id", "Hostname", "template", "Rplan", "Location", "Address", "Rplan"}
	printTable(header, data)
}

func scaletRemove(id int64) {
	client := vscale_api_go.NewClient(token)
	client.Scalet.Remove(id, false)
}
