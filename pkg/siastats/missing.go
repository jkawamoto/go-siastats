/*
 * missing.go
 *
 * Copyright (c) 2021 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package siastats

import (
	"fmt"
	"strconv"
	"strings"
)

type OneOfstringfloat struct {
	Value float32
}

func (s *OneOfstringfloat) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseFloat(strings.Trim(string(data), "\""), 32)
	if err != nil {
		return err
	}

	s.Value = float32(v)
	return nil
}

func (s *OneOfstringfloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(s.Value)), nil
}
