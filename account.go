package main

import (
	"github.com/vscale/go-vscale"
)

func getAccountInfo() {
	var data [][]string
	header := []string{"ID", "EMail", "Name"}
	client := vscale_api_go.NewClient(token)
	account, _, _ := client.Account.Get()
	data = append(data, []string{account.Info.ID, account.Info.Email, account.Info.Name + " " + account.Info.Surname})
	printTable(header, data)
}
