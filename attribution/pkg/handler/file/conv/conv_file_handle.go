/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 4:50 PM
 */

package conv

import (
	"fmt"

	"github.com/TencentAd/attribution/attribution/pkg/association"
	"github.com/TencentAd/attribution/attribution/pkg/association/validation"
	"github.com/TencentAd/attribution/attribution/pkg/handler/file/line"
	"github.com/TencentAd/attribution/attribution/pkg/parser"
	"github.com/TencentAd/attribution/attribution/pkg/parser/jsonline"
	"github.com/TencentAd/attribution/attribution/pkg/storage/attribution"
	"github.com/TencentAd/attribution/attribution/pkg/storage/clickindex"

	"github.com/golang/glog"
)

type FileHandle struct {
	filename         string
	parser           parser.ConvParserInterface
	clickIndex       clickindex.ClickIndex
	attributionStore []attribution.Storage
	assoc            *association.ClickAssociation
}

func NewConvFileHandle(filename string, clickIndex clickindex.ClickIndex, attributionStore []attribution.Storage) *FileHandle {
	assoc := association.NewClickAssociation().
		WithClickIndex(clickIndex).
		WithValidation(&validation.DefaultClickLogValidation{})
	return &FileHandle{
		filename:         filename,
		parser:           jsonline.NewConvParser(),
		clickIndex:       clickIndex,
		attributionStore: attributionStore,
		assoc:            assoc,
	}
}

func (p *FileHandle) Run() error {
	lp := line.NewLineProcess(p.filename, p.processLine, func(line string, err error) {
		glog.Errorf("failed to handle conv line[%s], err[%v]", line, err)
	}).WithParallelism(1)

	if err := lp.Run(); err != nil {
		return err
	}

	lp.WaitDone()
	return nil
}

func (p *FileHandle) processLine(line string) error {
	convLogs, err := p.parser.Parse(line)
	if err != nil {
		return err
	}

	for _, convLog := range convLogs.ConvLogs {
		c := &association.AssocContext{
			ConvLog: convLog,
		}

		if err := p.assoc.Association(c); err != nil {
			fmt.Println(err.Error())
			return err
		}

		for _, s := range p.attributionStore {
			if err = s.Store(c.ConvLog); err != nil {
				glog.Errorf("failed to store attribution result, err: %v", err)
			}
		}
	}

	return nil
}
