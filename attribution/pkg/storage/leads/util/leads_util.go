/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/11/20, 2:54 PM
 */

package util

import (
	"encoding/json"

	"github.com/TencentAd/attribution/attribution/pkg/leads/pull/protocal"
)

type LeadsIndexInfo struct {
	Key   string
	Value string
}

func ExtractLeadsIndex(leads *protocal.LeadsInfo) ([]*LeadsIndexInfo, error) {
	data, err := json.Marshal(leads)
	if err != nil {
		return nil, err
	}

	indexes := make([]*LeadsIndexInfo, 0)
	indexes = append(indexes, &LeadsIndexInfo{
		Key:   FormatKey(leads.LeadsTelephone),
		Value: string(data),
	})

	return indexes, nil
}

const prefix = "leads::"

func FormatKey(key string) string {
	return prefix + key
}
