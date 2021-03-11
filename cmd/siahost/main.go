/*
 * main.go
 *
 * Copyright (c) 2019-2021 Junpei Kawamoto
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

var (
	logger = log.New(os.Stderr, "", log.LstdFlags)
	client *siastats.APIClient
)

func init() {
	cfg := siastats.NewConfiguration()
	//cfg.Host = "https://siastats.info:3510/hosts-api"
	//cfg.Servers = [
	//]
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
		"allhosts",
		"summaries information and SiaStats final scores of all the hosts in the Sia network",
		"summaries information and SiaStats final scores of all the hosts in the Sia network", new(AllHostsCommand))
	if err != nil {
		logger.Fatalln("failed to add status command:", err)
	}

	_, err = parser.AddCommand(
		"host",
		"get information and detailed SiaStats benchmarks about a host. Uses the SiaStats-hostID",
		"get information and detailed SiaStats benchmarks about a host. Uses the SiaStats-hostID", new(HostCommand))
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
