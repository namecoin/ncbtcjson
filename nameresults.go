// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2019 The Namecoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package ncbtcjson

// NameShowResult models the data from the name_show command.
type NameShowResult struct {
	Name          string   `json:"name"`
	NameEncoding  Encoding `json:"name_encoding"`
	NameError     string   `json:"name_error"`
	Value         string   `json:"value"`
	ValueEncoding Encoding `json:"value_encoding"`
	ValueError    string   `json:"value_error"`
	TxID          string   `json:"txid"`
	Address       string   `json:"address"`
	Vout          uint32   `json:"vout"`
	Height        int32    `json:"height"`
	ExpiresIn     int32    `json:"expires_in"`
	IsMine        bool     `json:"ismine"`
	Expired       bool     `json:"expired"`
}

// NameScanResult models the data from the name_scan command.
type NameScanResult []NameShowResult
