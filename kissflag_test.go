// Package kissflag_test implements unit tests for the kissflag package
//
// Copyright 2018 William J House. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package kissflag_test

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/bhouse1273/kissflag"
)

const prefix string = "TEST_"

var resultStr string
var resultInt int
var resultInt32 int32
var resultInt64 int64
var resultBool bool
var resultFloat32 float32
var resultFloat64 float64

func TestBindEVar(t *testing.T) {
	type args struct {
		tag    string
		target interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		value   string
	}{
		{
			name: "Test string",
			args: args{
				tag:    "e1string",
				target: &resultStr,
			},
			wantErr: false,
			value:   "test1",
		},
		{
			name: "Test int",
			args: args{
				tag:    "e2int",
				target: &resultInt,
			},
			wantErr: false,
			value:   "1",
		},
		{
			name: "Test int32",
			args: args{
				tag:    "e3int32",
				target: &resultInt32,
			},
			wantErr: false,
			value:   "2",
		},
		{
			name: "Test bool",
			args: args{
				tag:    "e4bool",
				target: &resultBool,
			},
			wantErr: false,
			value:   "true",
		},
		{
			name: "Test float32",
			args: args{
				tag:    "e5float32",
				target: &resultFloat32,
			},
			wantErr: false,
			value:   "5.5",
		},
		{
			name: "Test float64",
			args: args{
				tag:    "e6float64",
				target: &resultFloat64,
			},
			wantErr: false,
			value:   "6.6",
		},
		{
			name: "Test empty string",
			args: args{
				tag:    "e7empty",
				target: &resultStr,
			},
			wantErr: false,
			value:   "",
		},
	}

	// Initialize prefix
	kissflag.SetPrefix(prefix)

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv(prefix+strings.ToUpper(tt.args.tag), tt.value)
			if err := kissflag.BindEVar(tt.args.tag, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("BindEVar() error = %v, wantErr %v", err, tt.wantErr)
			}
			switch tt.args.target.(type) {
			case *string:
				if *tt.args.target.(*string) != tt.value {
					err := errors.New("test value mismatch")
					t.Errorf("BindEVar() error = %v, wantErr %v", err, tt.wantErr)
				}
			case *int:
				if tval := strconv.FormatInt(int64(*tt.args.target.(*int)), 10); tval != tt.value {
					err := errors.New("test value mismatch")
					t.Errorf("BindEVar() error = %v, wantErr %v", err, tt.wantErr)
				}
			case *int32:
				if tval := strconv.FormatInt(int64(*tt.args.target.(*int32)), 10); tval != tt.value {
					err := errors.New("test value mismatch")
					t.Errorf("BindEVar() error = %v, wantErr %v", err, tt.wantErr)
				}
			case *int64:
				if tval := strconv.FormatInt(int64(*tt.args.target.(*int64)), 10); tval != tt.value {
					err := errors.New("test value mismatch")
					t.Errorf("BindEVar() error = %v, wantErr %v", err, tt.wantErr)
				}
			case *bool:
				if tval := strconv.FormatBool(*tt.args.target.(*bool)); tval != tt.value {
					err := errors.New("test value mismatch")
					t.Errorf("BindEVar() error = %v, wantErr %v", err, tt.wantErr)
				}

			case *float32:
				if tval := strconv.FormatFloat(float64(*tt.args.target.(*float32)), 'f', 1, 32); tval != tt.value {
					err := errors.New("test value mismatch")
					t.Errorf("BindEVar() error = %v, wantErr %v, value %v, result %v", err, tt.wantErr, tval, tt.value)
				}
			case *float64:
				if tval := strconv.FormatFloat(float64(*tt.args.target.(*float64)), 'f', 1, 64); tval != tt.value {
					err := errors.New("test value mismatch")
					t.Errorf("BindEVar() error = %v, wantErr %v, value %v, result %v", err, tt.wantErr, tval, tt.value)
				}
			}
		})
	}
}
