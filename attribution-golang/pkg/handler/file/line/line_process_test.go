/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/9/20, 5:26 PM
 */

package line

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

func printLineWithoutSpace(line string) error {
	if !strings.Contains(line, " ") {
		fmt.Println(line)
		return nil
	} else {
		return errors.New("contain space")
	}
}

func TestLineProcess(t *testing.T) {
	p := NewLineProcess("testdata/line.txt", printLineWithoutSpace,
		func(line string, err error) {
			fmt.Printf("line[%s] err: %v\n", line, err)
		}).
		WithParallelism(5).
		WithQueueSize(5).
		WithQueueTimeout(time.Millisecond * 10)

	p.Run()
	p.WaitDone()
}
