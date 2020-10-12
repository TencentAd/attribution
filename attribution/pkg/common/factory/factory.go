/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/12/20, 4:19 PM
 */

package factory

type Factory struct {
	Registered map[string]creator
}

func NewFactory() *Factory {
	return &Factory{
		Registered: make(map[string]creator),
	}
}

type creator func() interface{}

func (r *Factory) Register(name string, c creator) {
	r.Registered[name] = c
}

func (r *Factory) Create(name string) interface{} {
	if c, ok := r.Registered[name]; ok {
		return c()
	} else {
		return nil
	}
}
