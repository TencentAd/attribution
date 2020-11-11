/*
 * copyright (c) 2019, Tencent Inc.
 * All rights reserved.
 *
 * Author:  yuanweishi@tencent.com
 * Last Modify: 9/14/20, 4:26 PM
 */
package hbase

import (
	"sync"

	"github.com/tsuna/gohbase/hrpc"
	"github.com/tsuna/gohbase/pb"
)

type Row struct {
	R *hrpc.Result

	serializedData map[string]map[string][]byte
	serializeOnce  sync.Once
}

func (o *Row) serialize() {
	o.serializeOnce.Do(func() {
		o.serializedData = make(map[string]map[string][]byte)
		for _, c := range o.R.Cells {
			cell := (*pb.Cell)(c)
			family := string(cell.GetFamily())
			if _, ok := o.serializedData[family]; !ok {
				o.serializedData[family] = make(map[string][]byte)
			}
			o.serializedData[family][string(cell.GetQualifier())] = cell.GetValue()
		}
	})
}

func (o *Row) GetRowKey() string {
	cells := o.R.Cells
	if len(cells) == 0 {
		return ""
	}
	return string((*pb.Cell)(cells[0]).GetRow())
}
//func (o *Row) GetRowColumnQualifierValue() (string, string, string, string){
//	cells := o.R.Cells
//	if len(cells) == 0 {
//		return ""
//	}
//	return string((*pb.Cell)(cells[0]).GetRow())
//}

func (o *Row) GetValue(columnFamily, qualifier string) []byte {
	o.serialize()
	qualifiers, ok := o.serializedData[columnFamily]
	if !ok {
		return []byte{}
	}
	return qualifiers[qualifier]
}

func (o *Row) GetValueStr(columnFamily, qualifier string) string {
	return string(o.GetValue(columnFamily, qualifier))
}

type Column struct {
	RowKey       string
	ColumnFamily string
	Qualifier    string
	Value        []byte
}

func (o *Row) IterColumn() (ch chan *Column) {
	ch = make(chan *Column, 1)
	go func() {
		if o.R != nil {
			for _, c := range o.R.Cells {
				cell := (*pb.Cell)(c)
				ch <- &Column{
					string(cell.GetRow()),
					string(cell.GetFamily()),
					string(cell.GetQualifier()),
					cell.GetValue(),
				}
			}
		}
		close(ch)
	}()
	return
}










































