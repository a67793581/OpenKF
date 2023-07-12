// Copyright © 2023 OpenIM open source community. All rights reserved.
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

package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

const (
	_salt = "openkf"
)

// EncryptPassword encrypt password.
func EncryptPassword(password string) string {
	// md5
	m := md5.New()
	m.Write([]byte(password + _salt))
	mByte := m.Sum(nil)

	// hmac
	h := hmac.New(sha256.New, []byte(_salt))
	h.Write(mByte)
	password = hex.EncodeToString(h.Sum(nil))

	return password
}

// ComparePassword compare password.
func ComparePassword(password, encryptedPassword string) bool {
	return EncryptPassword(password) == encryptedPassword
}
