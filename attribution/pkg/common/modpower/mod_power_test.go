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
	prime = hex2Int(`C3DC4475E9AC81626966F699549130D3ABB461C33D98A7132AFD69A3C635EFEC0252C46ED722D52DF30BB6FBE5FAF38EBA3FFE09AFEE4939B9FB708C4ED3A803`)
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
	modPower1 := NewModPower(prime, big.NewInt(114343212421))
	assert.True(b, ok)

	for i :=0 ; i< b.N; i++ {
		modPower1.Encrypt(data)
	}
}

func hex2Int(hex string) *big.Int {
	data, _ := big.NewInt(0).SetString(hex, 16)
	return data
}