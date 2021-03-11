/*
 * hash.go
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
)

type HashCommand struct{}

func (c *HashCommand) Execute(args []string) error {
	for _, hash := range args {
		res, _, err := client.NavigatorApi.Hash(context.Background(), hash).Execute()
		if err != nil {
			return err
		}
		if len(res) < 2 {
			return fmt.Errorf("received result is not enough")
		}

		err = json.NewEncoder(os.Stdout).Encode(res[1])
		if err != nil {
			return err
		}

	}

	return nil
}
