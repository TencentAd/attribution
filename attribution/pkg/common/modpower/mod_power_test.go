/*
 * copyright (c) 2020, Tencent Inc.
 * All rights reserved.
 *
 * Author:  linceyou@tencent.com
 * Last Modify: 10/23/20, 2:52 PM
 */

package modpower

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	prime = MustHex2BigInt(`CC5A0467495EC2506D699C9FFCE6ED5885AE09384671EE00B93A28712DA814240F2A471B2C77B120FE70DA02D33B0F85CD737B800942D5A2BF80DCCD290FB6553C9834197DC498F6AC69B5ECEF3FD8B05F231E3632E9AECA2B3F50977CF9033AF3005A9C0A339CFB4922971B3AF05A5955983C12B153BB78A2B1FB14C84A3C662ADDE5BCEBE8779FF9F97C6E73BD29D4044242581455EEB0543E2DB35F43997F46F8596A58080DC053BBB71F9A557185DF80738238713D3EDFD77D47B26977B373FB0D969920B3909CCC24792B5B4E94AD29F6AE6BD73ED5FE6528CDDBEA1560BBCD36E8B25008021A26E9A4E51BBCD8436F38D6A222E2138E7042A73A7877D7`)
)

func TestModPower_Decrypt(t *testing.T) {
	originalText := `797332e3b842d5b5d4441e1278362e2f`
	data, ok := big.NewInt(0).SetString(originalText, 16)
	assert.True(t, ok)

	modPower := NewModPower(prime, big.NewInt(1143421))
	encData := modPower.Encrypt(data)
	t.Log(encData.Text(16))
	decData := modPower.Decrypt(encData)
	assert.EqualValues(t, originalText, decData.Text(16))
}

func TestModPower_Exchange(t *testing.T) {
	originalText := `797332e3b842d5b5d4441e1278362e2f`
	data, ok := big.NewInt(0).SetString(originalText, 16)
	assert.True(t, ok)

	e1 := big.NewInt(1143421)
	e2 := big.NewInt(114343212421)

	modPower1 := NewModPower(prime, e1)
	modPower2 := NewModPower(prime, e2)

	{
		enc1Data := modPower1.Encrypt(data)
		enc12Data := modPower2.Encrypt(enc1Data)

		enc2Data := modPower2.Encrypt(data)
		enc21Data := modPower1.Encrypt(enc2Data)

		assert.EqualValues(t, enc12Data.String(), enc21Data.String())
	}

	{
		enc1Data := modPower1.Encrypt(data)
		enc12Data := modPower2.Encrypt(enc1Data)
		enc12Dec1Data := modPower1.Decrypt(enc12Data)

		enc2Data := modPower2.Encrypt(data)

		assert.EqualValues(t, enc12Dec1Data.String(), enc2Data.String())
	}
}

func BenchmarkNewModPower(b *testing.B) {
	originalText := `277cf7d3d25be8165aba097ff4c599e001a316b2654acd757ca80b1ea56012a7`
	data, ok := big.NewInt(0).SetString(originalText, 16)
	modPower1 := NewModPower(prime, MustHex2BigInt("9cff5b8e1899bb3e7a356f3eb6b45a85"))
	assert.True(b, ok)

	for i :=0 ; i< b.N; i++ {
		modPower1.Encrypt(data)
	}
}

func hex2Int(hex string) *big.Int {
	data, _ := big.NewInt(0).SetString(hex, 16)
	return data
}