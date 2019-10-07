/*
 * main.go
 *
 * Copyright (c) 2019 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
	siastats "github.com/jkawamoto/go-siastats/pkg"
)

var client *siastats.APIClient
var logger = log.New(os.Stderr, "", log.LstdFlags)

func init() {
	cfg := siastats.NewConfiguration()
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client = siastats.NewAPIClient(cfg)
}

func main() {

	parser := flags.NewParser(nil, flags.Default)
	_, err := parser.AddCommand(
		"status",
		"get status and blockchain sync situation of SiaStats nodes",
		"get status and blockchain sync situation of SiaStats nodes", new(StatusCommand))
	if err != nil {
		logger.Fatalln("failed to add status command:", err)
	}

	_, err = parser.AddCommand(
		"hash",
		"returns the information about any hash or block height on the blockchain",
		"returns the information about any hash (address, Tx, contract...) or block height on the blockchain", new(HashCommand))
	if err != nil {
		logger.Fatalln("failed to add status command:", err)
	}

	_, err = parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

}
