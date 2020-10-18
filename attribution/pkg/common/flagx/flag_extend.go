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
	"io/ioutil"
	"os"
	"strings"
)

type flagFileSet map[string]struct{}

func (a *flagFileSet) Set(value string) error {
	if _, ok := (*a)[value]; !ok {
		(*a)[value] = struct{}{}
	}
	return nil
}

func (a *flagFileSet) String() string {
	flagFileList := make([]string, 0, len(*a))
	for key := range *a {
		flagFileList = append(flagFileList, key)
	}
	return strings.Join(flagFileList, ",")
}

var (
	FlagFiles flagFileSet = make(map[string]struct{})

	ignoreUndefinedFlag = flag.Bool("ignore_undefined_flag", false, "")
)

func init() {
	flag.Var(&FlagFiles, "flagfile", "")
}

func printFlag(f *flag.Flag) {
	fmt.Println("-" + f.Name + "=" + f.Value.String())
}

func replaceWithEnv(f *flag.Flag) {
	if value, ok := os.LookupEnv(f.Name); ok {
		f.Value.Set(value)
	}
}

// 自带的flag库不支持传入文件，这个接口支持传入"-flagfile"方便传入配置文件
func Parse() error {
	flag.Parse()
	// 1. 读取环境变量
	flag.VisitAll(replaceWithEnv)

	// 2. 读取flagfile
	var validFlagLines []string
	for f := range FlagFiles {
		flagContents, err := ioutil.ReadFile(f)
		if err != nil {
			return fmt.Errorf("err when read content from flagfile[%s], err:%s", f, err)
		}
		flagLines := strings.Split(string(flagContents), "\n")
		for _, line := range flagLines {
			if len(line) != 0 && string([]rune(line)[0]) != "#" {
				if *ignoreUndefinedFlag {
					v := strings.Split(line, "=")
					if len(v) > 0 {
						n := v[0]
						n = strings.TrimLeft(n, " \t-")
						n = strings.TrimRight(n, " \t")
						if flag.CommandLine.Lookup(n) == nil {
							continue
						}
					}
				}
				validFlagLines = append(validFlagLines, line)
			}
		}
	}
	err := flag.CommandLine.Parse(validFlagLines)
	if err != nil {
		return fmt.Errorf("failed to parse flagfile[%s], err:%s", FlagFiles.String(), err)
	}

	// 3. 主动在命令行设置的参数
	flag.Parse()
	flag.VisitAll(printFlag)

	return nil
}
