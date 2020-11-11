/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 2:31 PM
 */

package define

import "flag"

const (
	VLogLevel = 10
)

var (
	LeadsExpireHour = flag.Int("leads_expire_hour", 720, "")
)