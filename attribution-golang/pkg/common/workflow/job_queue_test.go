/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:28 AM
 */

package workflow

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultJobQueue_ReceiveJob(t *testing.T) {
	q := NewDefaultJobQueue(
		&QueueOption{
			WorkerCount: 2,
			QueueSize:   2,
			PushTimeout: time.Second,
		})

	q.Start()

	for i := 0; i < 10; i++ {
		{
			k := i
			err := q.PushJob(func() {
				fmt.Println("i = ", k)
			})

			assert.NoError(t, err)
		}
	}

	q.Stop()
	time.Sleep(time.Second)
}
