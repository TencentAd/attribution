/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/12/20, 5:35 PM
 */

package click

import (
	"github.com/TencentAd/attribution/attribution/pkg/handler/file/line"
	"github.com/TencentAd/attribution/attribution/pkg/logic"
	"github.com/TencentAd/attribution/attribution/pkg/parser"
	"github.com/TencentAd/attribution/attribution/pkg/parser/jsonline"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"

	"github.com/golang/glog"
)

type FileHandle struct {
	parser     parser.ClickParserInterface
	clickIndex clickindex.ClickIndex
	filename   string
}

func NewClickFileHandle(filename string, clickIndex clickindex.ClickIndex) *FileHandle {
	return &FileHandle{
		parser:     jsonline.NewJsonClickParser(),
		filename:   filename,
		clickIndex: clickIndex,
	}
}

func (p *FileHandle) Run() error {
	lp := line.NewLineProcess(p.filename, p.processLine, func(line string, err error) {
		glog.Errorf("failed to handle line[%s], err[%v]", line, err)
	})

	if err := lp.Run(); err != nil {
		return err
	}

	lp.WaitDone()
	return nil
}

func (p *FileHandle) processLine(line string) error {
	clickLog, err := p.parser.Parse(line)
	if err != nil {
		return err
	}
	return logic.ProcessClickLog(clickLog, p.clickIndex)
}
