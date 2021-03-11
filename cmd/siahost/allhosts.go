/*
 * allhosts.go
 *
 * Copyright (c) 2019-2021 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package main

import (
	"context"
	"encoding/json"
	"os"

	siastats "github.com/jkawamoto/go-siastats/pkg"
)

type AllHostsCommand struct{}

func (c *AllHostsCommand) Execute(_ []string) (err error) {
	hosts, _, err := client.HostApi.AllHosts(context.Background()).AllHostsRequest(siastats.AllHostsRequest{
		Network: "sia",
		List:    "active",
	}).Execute()
	if err != nil {
		return
	}

	return json.NewEncoder(os.Stdout).Encode(hosts)
}
