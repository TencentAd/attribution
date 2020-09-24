/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/12/20, 5:35 PM
 */

package click

import (
	"attribution/pkg/logic"
	"attribution/pkg/storage"
	"github.com/golang/glog"
	"attribution/pkg/parser"
	"attribution/pkg/parser/click_parser"
)

type ClickFileHandle struct {
	parser parser.ClickParserInterface
	clickIndex storage.ClickIndex
	filename   string
}

func NewClickFileHandle(filename string, clickIndex storage.ClickIndex) *ClickFileHandle {
	return &ClickFileHandle{
		parser:     click_parser.NewClickParser(),
		filename:   filename,
		clickIndex: clickIndex,
	}
}

func (p *ClickFileHandle) Run() error {
	lp := NewLineProcess(p.filename, p.processLine, func(line string, err error) {
		glog.Errorf("failed to handle line[%s], err[%v]", line, err)
	})

	if err := lp.Run(); err != nil {
		return err
	}

	lp.WaitDone()
	return nil
}

func (p *ClickFileHandle) processLine(line string) error {
	clickLog, err := p.parser.Parse(line)
	if err != nil {
		return err
	}
	return logic.ProcessClickLog(clickLog, p.clickIndex)
}
