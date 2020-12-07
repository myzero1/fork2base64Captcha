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

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/mojocn/base64Captcha"
)

func TestNoStoreCaptcha_GenerateB64s_Verify(t *testing.T) {
	store := base64Captcha.NewNoStore("mzyero1", 30, 10)
	bgc := &color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	driver := base64Captcha.NewDriverString(60,
		180,
		0,
		14, // 2/4/8
		6,
		"1234567890qwertyuioplkjhgfdsazxcvbnm",
		bgc,
		[]string{})

	c := base64Captcha.NewNoStoreCaptcha(driver, store)
	id, b64s, a, err := c.Generate()

	if err != nil {
		t.Errorf("NoStoreCaptcha.Generate() error = %v", err)
		return
	}

	fmt.Println("id and answer are:", id, a)

	t.Log(b64s)

	if !c.Verify(id, a, true) {
		t.Error("false")
	}
}

func TestNoStoreCaptcha_GenerateIdQuestionAnswer_Verify(t *testing.T) {
	store := base64Captcha.NewNoStore("myzero1", 30, 10)
	bgc := &color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	driver := base64Captcha.NewDriverString(60,
		180,
		0,
		14, // 2/4/8
		6,
		"1234567890qwertyuioplkjhgfdsazxcvbnm",
		bgc,
		[]string{})

	c := base64Captcha.NewNoStoreCaptcha(driver, store)
	id, content, answer := c.GenerateIdQuestionAnswer()

	fmt.Println("id,content and answer are:", id, content, answer)

	if !c.Verify(id, answer, true) {
		t.Error("false")
	}
}
