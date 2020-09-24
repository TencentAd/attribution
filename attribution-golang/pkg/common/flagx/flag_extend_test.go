/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 9/9/20, 9:30 AM
 */
package flagx

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	testString       string
	testQuotedString string
	testInt          int
)

func init() {
	flag.StringVar(&testString, "test_string", "", "")
	flag.StringVar(&testQuotedString, "test_quoted_string", "", "")
	flag.IntVar(&testInt, "test_int", 0, "")
}

func TestParse(t *testing.T) {
	os.Args = []string{"null", "-flagfile=testdata/test.flags", "--flagfile=testdata/test2.flags"}
	if err := Parse(); err != nil {
		t.Error(err)
	}

	if testString != "string" {
		t.Errorf("failed to read string flag")
	}

	if testQuotedString != `"quoted_string"` {
		t.Errorf("failed to read quoted string flag")
	}

	if testInt != 1 {
		t.Errorf("failed to read int flag")
	}
}

func TestParseFromCommandLine(t *testing.T) {
	os.Args = []string{"null", "-test_int=1"}
	if err := Parse(); err != nil {
		t.Error(err)
	}

	if testInt != 1 {
		t.Errorf("failed to read int flag")
	}
}

func TestIgnoreUndefinedFlag(t *testing.T) {
	os.Args = []string{"dummy", "-flagfile=testdata/test3.flags", "-ignore_undefined_flag=true"}

	fmt.Printf("@@@start testing ignore undefined flag:%+v@@@@\n", os.Args)

	for k := range FlagFiles {
		delete(FlagFiles, k)
	}

	if err := Parse(); err != nil {
		t.Errorf("@@@@ ut failed:%s", err)
	}

	fmt.Printf("@@@@@ignore undefine flag:%v", *ignoreUndefinedFlag)
}
