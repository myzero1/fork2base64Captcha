// Copyright 2017 Eric Zhou. All Rights Reserved.
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

package base64Captcha

// noStore is an internal store for captcha ids and their values.
type noStore struct {
	// unit is Second
	Timeout int
}

// NewNoStore returns a new standard memory store for captchas with the
// given collection threshold and expiration time (duration). The returned
// store must be registered with SetCustomStore to replace the default one.
func NewNoStore(timeout ...int) *noStore {
	s := new(noStore)
	s.Timeout = 300
	if len(timeout) > 0 {
		s.Timeout = timeout[0]
	}
	return s
}

func (s *noStore) Set(id string, value string) {
}

func (s *noStore) Verify(id, answer string, clear bool) (ok bool) {
	md5Id := generateMD5ID(answer, s.Timeout)

	return md5Id == id
}

func (s *noStore) Get(id string, clear bool) (value string) {
	return "myzero1"
}
