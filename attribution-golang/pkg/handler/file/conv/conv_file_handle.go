/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 8/13/20, 4:50 PM
 */

package conv

import (
	"encoding/json"
	"fmt"

	"attribution/pkg/association"
	"attribution/pkg/association/validation"
	"attribution/pkg/handler/file/line"
	"attribution/pkg/parser"
	"attribution/pkg/parser/conv_parser"
	"attribution/pkg/storage"

	"github.com/golang/glog"
)

type FileHandle struct {
	filename         string
	parser           parser.ConvParserInterface
	clickIndex       storage.ClickIndex
	attributionStore storage.AttributionStore
	assoc            *association.ClickAssociation
}

func NewConvFileHandle(filename string, clickIndex storage.ClickIndex, attributionStore storage.AttributionStore) *FileHandle {
	assoc := association.NewClickAssociation().
		WithClickIndex(clickIndex).
		WithValidation(&validation.DefaultClickLogValidation{})
	return &FileHandle{
		filename:         filename,
		parser:           conv_parser.NewConvParser(),
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
	convLog, err := p.parser.Parse(line)
	if err != nil {
		return err
	}

	c := &association.AssocContext{
		ConvLog: convLog,
	}

	if err := p.assoc.Association(c); err != nil {
		fmt.Println(err.Error())
		return err
	}

	p.attributionStore.Store(c.ConvLog, c.SelectClick)
	return nil
}

func printAssociationResult(c *association.AssocContext) error {
	var convContent []byte
	var clickContent []byte
	var err error
	convContent, err = json.Marshal(c.ConvLog)
	if err != nil {
		return err
	}

	if c.SelectClick != nil {
		clickContent, err = json.Marshal(c.SelectClick)
		if err != nil {
			return err
		}
	}

	fmt.Printf("conv: %s\nclick:%s\n", string(convContent), string(clickContent))
	return nil
}
