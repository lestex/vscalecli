package main

import (
	"strconv"

	"github.com/vscale/go-vscale"
)

// (makeFrom, rplan, name, password, location string, doStart bool,keys []int64, wait bool)
func createScalet(from, rplan, name, location string, doStart bool, sshkeys string) *vscale_api_go.Scalet {
	client := vscale_api_go.NewClient(token)
	keys := stringToKeys(sshkeys)
	scalet, _, _ := client.Scalet.CreateWithoutPassword(from,
		rplan,
		name,
		location,
		doStart,
		keys,
		false)
	return scalet
}

func listScalet() {
	client := vscale_api_go.NewClient(token)
	scalets, _, _ := client.Scalet.List()
	var data [][]string

	for _, s := range *scalets {
		data = append(data, []string{strconv.Itoa(int(s.CTID)),
			s.Hostname, s.MadeFrom, s.Rplan, s.Location, s.PublicAddresses.Address, s.Status})
	}
	header := []string{"id",
		"Hostname",
		"template",
		"Rplan",
		"Location",
		"Address",
		"status"}
	printTable(header, data)
}

func removeScalet(id int64) {
	client := vscale_api_go.NewClient(token)
	client.Scalet.Remove(id, false)
}

func restartScalet(id int64) {
	client := vscale_api_go.NewClient(token)
	client.Scalet.Restart(id, false)
}

func startScalet(id int64) {
	client := vscale_api_go.NewClient(token)
	client.Scalet.Start(id, false)
}

func stopScalet(id int64) {
	client := vscale_api_go.NewClient(token)
	client.Scalet.Stop(id, false)
}
