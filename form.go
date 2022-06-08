// Copyright 2014 Alvaro J. Genial. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package form implements encoding and decoding of application/x-www-form-urlencoded data.
package form

import (
	"net/url"
	"strings"
)

const (
	implicitKey = "_"
	omittedKey  = "-"

	defaultDelimiter = '.'
	defaultEscape    = '\\'
)

func TrancferToParentheses(values url.Values) url.Values {
	newValues := url.Values{}

	for key, value := range values {
		keys := strings.Split(key, ".")

		var newKeys string

		if len(keys) > 1 {
			newKeys = keys[0]
			for _, k := range keys[1:] {
				newKeys += "[" + k + "]"
			}
		} else {
			newKeys = strings.Join(keys, "")
		}
		for _, v := range value {
			newValues.Add(newKeys, v)
		}
	}

	return newValues
}
