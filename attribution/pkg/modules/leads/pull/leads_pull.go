/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/25/20, 10:22 AM
 */

package main

import (
	"flag"

	"github.com/TencentAd/attribution/attribution/pkg/common/flagx"
	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/client"
	"github.com/TencentAd/attribution/attribution/pkg/storage"
)

var (
	leadsPullConfig = flag.String("leads_pull_config", "", "")
)

func run() error {
	pullConfig, err := client.NewPullConfig(*leadsPullConfig)
	if err != nil {
		return err
	}

	leadsStorage, err := storage.CreateLeadsStorage()
	if err != nil {
		return err
	}

	pullClient := client.NewLeadsPullClient(pullConfig).WithStorage(leadsStorage)
	return pullClient.Pull()
}

func main() {
	err := flagx.Parse()
	if err != nil {
		panic(err)
	}

	if err := run(); err != nil {
		panic(err)
	}
}
