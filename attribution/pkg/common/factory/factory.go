/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/12/20, 4:19 PM
 */

package factory

import "fmt"

type Factory struct {
	factoryName string
	Registered  map[string]creator
}

func NewFactory(factoryName string) *Factory {
	return &Factory{
		factoryName: factoryName,
		Registered:  make(map[string]creator),
	}
}

type creator func() interface{}

func (r *Factory) Register(name string, c creator) {
	r.Registered[name] = c
}

func (r *Factory) Create(name string) (interface{}, error) {
	if c, ok := r.Registered[name]; ok {
		if obj :=  c(); obj == nil {
			return nil, fmt.Errorf("failed to create factory[%s] name[%s]", r.factoryName, name)
		} else {
			return obj, nil
		}
	} else {
		return nil, fmt.Errorf("factory[%s] not support [%s]", r.factoryName, name)
	}
}
