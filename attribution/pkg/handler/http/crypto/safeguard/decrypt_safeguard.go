/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 11/10/20, 2:54 PM
 */

package safeguard

type DecryptSafeguard struct {

}

func NewDecryptSafeguard() *DecryptSafeguard {
	return &DecryptSafeguard{}
}

// TODO
func (g *DecryptSafeguard) Against(opt *Parameter) error {
	return nil
}
