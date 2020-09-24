/*
 * Copyright (c) 2018-2118
 * Author: linceyou
 * LastModified: 19-4-29 下午3:33
 */

package metricUtil

import (
	"time"
)

type Observed interface{
	labels() []string
}

func CalcTimeUsedMicro(startTime time.Time) float64 {
	return float64(time.Since(startTime) / time.Microsecond)
}

func CalcTimeUsedMilli(startTime time.Time) float64 {
	return float64(time.Since(startTime) / time.Millisecond)
}