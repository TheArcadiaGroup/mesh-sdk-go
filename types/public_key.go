// Copyright 2024 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

import (
	"encoding/hex"
	"encoding/json"
)

// PublicKey PublicKey contains a public key byte array for a particular CurveType encoded in hex.
// Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
type PublicKey struct {
	Bytes     []byte    `json:"hex_bytes"`
	CurveType CurveType `json:"curve_type"`
}

// MarshalJSON overrides the default JSON marshaler
// and encodes bytes as hex instead of base64.
func (s *PublicKey) MarshalJSON() ([]byte, error) {
	type Alias PublicKey
	j, err := json.Marshal(struct {
		Bytes string `json:"hex_bytes"`
		*Alias
	}{
		Bytes: hex.EncodeToString(s.Bytes),
		Alias: (*Alias)(s),
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// UnmarshalJSON overrides the default JSON unmarshaler
// and decodes bytes from hex instead of base64.
func (s *PublicKey) UnmarshalJSON(b []byte) error {
	type Alias PublicKey
	r := struct {
		Bytes string `json:"hex_bytes"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	bytes, err := hex.DecodeString(r.Bytes)
	if err != nil {
		return err
	}

	s.Bytes = bytes
	return nil
}
