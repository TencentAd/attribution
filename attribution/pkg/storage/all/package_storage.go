/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/12/20, 5:06 PM
 */

package all

import (
	_ "github.com/TencentAd/attribution/attribution/pkg/storage/hbase"
	_ "github.com/TencentAd/attribution/attribution/pkg/storage/hdfs"
	_ "github.com/TencentAd/attribution/attribution/pkg/storage/http"
	_ "github.com/TencentAd/attribution/attribution/pkg/storage/native"
	_ "github.com/TencentAd/attribution/attribution/pkg/storage/redis"
)
