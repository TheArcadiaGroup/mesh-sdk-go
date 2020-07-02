// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keys

import (
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateKeypairSecp256k1(t *testing.T) {
	curve := types.Secp256k1
	keypair, err := GenerateKeypair(curve)

	assert.NoError(t, err)
	assert.Equal(t, keypair.PublicKey.CurveType, curve)
	assert.Equal(t, keypair.PrivateKey.CurveType, curve)
}

func TestGenerateKeypairEd25519(t *testing.T) {
	curve := types.Edwards25519
	keypair, err := GenerateKeypair(curve)

	assert.NoError(t, err)
	assert.Equal(t, keypair.PublicKey.CurveType, curve)
	assert.Equal(t, keypair.PrivateKey.CurveType, curve)
}

type privKeyTest struct {
	keypair *KeyPair
	errMsg  string
}

func mockKeyPair(privKey []byte, curveType types.CurveType) *KeyPair {
	keypair, _ := GenerateKeypair(curveType)
	keypair.PrivateKey.Bytes = privKey
	return keypair
}

var privKeyTests = []privKeyTest{
	{mockKeyPair(make([]byte, 33), types.Secp256k1), "invalid privkey length"},
	{mockKeyPair(make([]byte, 31), types.Secp256k1), "invalid privkey length"},
	{mockKeyPair(make([]byte, 0), types.Secp256k1), "invalid privkey length"},
	{mockKeyPair(make([]byte, 33), types.Edwards25519), "invalid privkey length"},
	{mockKeyPair(make([]byte, 31), types.Edwards25519), "invalid privkey length"},
	{mockKeyPair(make([]byte, 0), types.Edwards25519), "invalid privkey length"},
}

func TestKeypairValidity(t *testing.T) {
	// Non matching curves
	keyPair, _ := GenerateKeypair(types.Edwards25519)
	keyPair.PublicKey.CurveType = types.Secp256k1
	err := keyPair.IsValid()
	assert.Contains(t, err.Error(), "do not match")

	for _, test := range privKeyTests {
		err = test.keypair.IsValid()
		assert.Contains(t, err.Error(), test.errMsg)
	}
}
