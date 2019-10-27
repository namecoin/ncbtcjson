// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2019 The Namecoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: This file is intended to house the RPC commands that are supported by
// a name lookup client.

package ncbtcjson

import (
	"github.com/namecoin/btcd/btcjson"
)

// Encoding represents an encoding for name identifiers and values.
type Encoding string

const (
	Ascii Encoding = "ascii"
	Utf8           = "utf8"
	Hex            = "hex"
)

// NameShowOptions represents the optional options struct provided with a
// NameShowCmd command.
type NameShowOptions struct {
	NameEncoding  Encoding `json:"nameEncoding,omitempty"`
	ValueEncoding Encoding `json:"valueEncoding,omitempty"`
	StreamID      string   `json:"streamID,omitempty"`
}

// NameShowCmd defines the name_show JSON-RPC command.
type NameShowCmd struct {
	Name    string
	Options *NameShowOptions
}

// NewNameShowCmd returns a new instance which can be used to issue a
// name_show JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewNameShowCmd(name string, options *NameShowOptions) *NameShowCmd {
	return &NameShowCmd{
		Name:    name,
		Options: options,
	}
}

// NameScanOptions represents the optional options struct provided with a
// NameScanCmd command.
type NameScanOptions struct {
	NameEncoding  Encoding `json:"nameEncoding,omitempty"`
	ValueEncoding Encoding `json:"valueEncoding,omitempty"`
	MinConf       int32    `json:"minConf,omitempty"`
	MaxConf       int32    `json:"maxConf,omitempty"`
	Prefix        string   `json:"prefix,omitempty"`
	RegExp        string   `json:"regexp,omitempty"`
}

// NameScanCmd defines the name_scan JSON-RPC command.
type NameScanCmd struct {
	Start   string
	Count   *uint32
	Options *NameScanOptions
}

// NewNameScanCmd returns a new instance which can be used to issue a
// name_scan JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewNameScanCmd(start string, count *uint32, options *NameScanOptions) *NameScanCmd {
	return &NameScanCmd{
		Start:   start,
		Count:   count,
		Options: options,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := btcjson.UsageFlag(0)

	btcjson.MustRegisterCmd("name_show", (*NameShowCmd)(nil), flags)
	btcjson.MustRegisterCmd("name_scan", (*NameScanCmd)(nil), flags)
}
