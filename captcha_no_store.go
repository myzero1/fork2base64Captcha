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

// Package base64Captcha supports digits, numbers,alphabet, arithmetic, audio and digit-alphabet captcha.
// base64Captcha is used for fast development of RESTful APIs, web apps and backend services in Go. give a string identifier to the package and it returns with a base64-encoding-png-string
package base64Captcha

// NoStoreCaptcha captcha basic information.
type NoStoreCaptcha struct {
	Driver Driver
	Store  *noStore
}

//NewNoStoreCaptcha creates a captcha instance from driver and store
func NewNoStoreCaptcha(driver Driver, store *noStore) *NoStoreCaptcha {
	return &NoStoreCaptcha{Driver: driver, Store: store}
}

//Generate generates a random id, base64 image string or an error if any
func (c *NoStoreCaptcha) Generate() (id, b64s, answer string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		return "", "", "", err
	}

	// c.Store.Set(id, answer)

	// timeout := 300
	// idLength := 8
	// myValues := reflect.ValueOf(*c.Store)
	// myTypes := myValues.Type()
	// for i := 0; i < myValues.NumField(); i++ {
	// 	fieldName := myTypes.Field(i).Name
	// 	if "Timeout" == fieldName {
	// 		tmp, err := strconv.Atoi(fmt.Sprintf(`%v`, myValues.Field(i)))
	// 		if err == nil {
	// 			timeout = tmp
	// 		}
	// 	} else {
	// 		if "IDLength" == fieldName {
	// 			tmp, err := strconv.Atoi(fmt.Sprintf(`%v`, myValues.Field(i)))
	// 			if err == nil {
	// 				idLength = tmp
	// 			}
	// 		}
	// 	}
	// }

	// id = GenerateMD5ID(answer, c.Store.Password, timeout, idLength)

	id = GenerateMD5ID(answer, c.Store.Password, c.Store.Timeout, c.Store.IDLength)

	b64s = item.EncodeB64string()
	return
}

//Verify by a given id key and remove the captcha value in store,
//return boolean value.
//if you has multiple captcha instances which share a same store.
//You may want to call `store.Verify` method instead.
func (c *NoStoreCaptcha) Verify(id, answer string, clear bool) (match bool) {
	return c.Store.Verify(id, answer, clear)
}

//GenerateIdQuestionAnswer creates id,content and answer
func (c *NoStoreCaptcha) GenerateIdQuestionAnswer() (id, content, answer string) {
	id, content, answer = c.Driver.GenerateIdQuestionAnswer()
	id = GenerateMD5ID(answer, c.Store.Password, c.Store.Timeout, c.Store.IDLength)
	return id, content, content
}
