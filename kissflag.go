// Package kissflag implements KISS flag helper for env vars
//
// Copyright 2018 William J House. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package kissflag

import (
	"encoding/base64"
	"errors"
	"os"
	"strconv"
	"strings"
)

var prefix string

// BindEVar assigns the value of a named configuration value
func BindEVar(tag string, target interface{}) error {
	var err error
	// Enforce non-empty tag value
	if tag == "" {
		err = errors.New("tag argument may not be empty")
		return err
	}
	// Enforce non-nil target
	if target == nil {
		err = errors.New("target argument may not be nil")
		return err
	}
	// Enforce naming convention assuming a prefix match indicates tag is preformatted
	if len(tag) <= len(prefix) || strings.ToUpper(tag[0:len(prefix)]) != prefix {
		tag = prefix + strings.ToUpper(tag)
	}
	// Lookup tag value in environment allowing an empty value to be assigned if present
	if value, ok := os.LookupEnv(tag); ok {
		switch target.(type) {
		case *string:
			*target.(*string) = value
		case *bool:
			*target.(*bool), err = strconv.ParseBool(value)
		case *int:
			*target.(*int), err = strconv.Atoi(value)
		case *int32:
			t := int64(0)
			if t, err = strconv.ParseInt(value, 10, 32); err == nil {
				*target.(*int32) = int32(t)
			}
		case *int64:
			*target.(*int64), err = strconv.ParseInt(value, 10, 64)
		case *float32:
			t := float64(0.0)
			if t, err = strconv.ParseFloat(value, 32); err == nil {
				*target.(*float32) = float32(t)
			}
		case *float64:
			*target.(*float64), err = strconv.ParseFloat(value, 64)
		default:
			err = errors.New("Unsupported target type")
		}
	}
	return err
}

// DecodeBase64 attempts to decode value and sets the target if successful,
// returns error if not. Note that some values may give false positives if
// size is not provided (i.e., size == 0)
func DecodeBase64(value string, target *string, size int) error {
	var err error
	var tval []byte
	if tval, err = base64.StdEncoding.DecodeString(value); err == nil {
		if size > 0 && len(string(tval)) != size {
			err = errors.New("target size mismatch")
			return err
		}
		*target = string(tval)
	}
	return err
}

// SetPrefix assigns the value of the hidden prefix variable
func SetPrefix(value string) {
	prefix = value
}
