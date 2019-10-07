/*
 * status.go
 *
 * Copyright (c) 2019 Junpei Kawamoto
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
)

type StatusCommand struct{}

func (s *StatusCommand) Execute(args []string) (err error) {

	res, _, err := client.NavigatorApi.Status(context.Background())
	if err != nil {
		return
	} else if len(res) == 0 {
		return fmt.Errorf("got an empty result")
	}

	err = json.NewEncoder(os.Stdout).Encode(res[0])
	return

}
