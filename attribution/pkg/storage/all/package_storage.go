/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/12/20, 5:06 PM
 */

package all

import (
	_ "attribution/pkg/storage/hbase"
	_ "attribution/pkg/storage/hdfs"
	_ "attribution/pkg/storage/http"
	_ "attribution/pkg/storage/native"
	_ "attribution/pkg/storage/redis"
)
