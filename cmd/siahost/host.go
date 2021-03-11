/*
 * host.go
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
	"fmt"
	"os"
	"strconv"
)

type HostCommand struct {
}

func (c *HostCommand) Execute(args []string) error {
	ctx := context.Background()
	for _, v := range args {

		id, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("failed to convert given host ID %v: %v", v, err)
		}

		res, _, err := client.HostApi.Host(ctx, int32(id)).Execute()
		if err != nil {
			return fmt.Errorf("failed to get information of host %v: %v", id, err)
		}

		err = json.NewEncoder(os.Stdout).Encode(res)
		if err != nil {
			return fmt.Errorf("failed to write host information: %v", err)
		}

	}

	return nil
}
