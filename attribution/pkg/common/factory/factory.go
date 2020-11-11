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
	factoryName             string
	registered              map[string]creator
	optionCreatorRegistered map[string]optionCreator
}

func NewFactory(factoryName string) *Factory {
	return &Factory{
		factoryName:             factoryName,
		registered:              make(map[string]creator),
		optionCreatorRegistered: make(map[string]optionCreator),
	}
}

type creator func() interface{}

func (r *Factory) Register(name string, c creator) {
	r.registered[name] = c
}

type optionCreator func(option interface{}) interface{}

func (r *Factory) RegisterOptionCreator(name string, c optionCreator) {
	r.optionCreatorRegistered[name] = c
}

func (r *Factory) Create(name string) (interface{}, error) {
	if c, ok := r.registered[name]; ok {
		if obj := c(); obj == nil {
			return nil, fmt.Errorf("failed to create factory[%s] name[%s]", r.factoryName, name)
		} else {
			return obj, nil
		}
	} else {
		return nil, fmt.Errorf("factory[%s] not support [%s]", r.factoryName, name)
	}
}

func (r *Factory) CreateWithOption(name string, option interface{}) (interface{}, error) {
	if c, ok := r.optionCreatorRegistered[name]; ok {
		if obj := c(option); obj == nil {
			return nil, fmt.Errorf("failed to create factory[%s] name[%s]", r.factoryName, name)
		} else {
			return obj, nil
		}
	} else {
		return nil, fmt.Errorf("factory[%s] not support [%s]", r.factoryName, name)
	}
}
