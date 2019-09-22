// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf8E\u0594\x04\xee\u040cB\x00\xcd\xf5\x14\xa9\xb0y\xcff\x9bqr\xe7\x1d\xfa\x80\u0594\x99o)Y\xcehK,\xa2!\xb9\xf0\xdaA\x89\x96b\"\tS\x80\u0594\xe0\xc1E\xcb\xe3A{r\xa1\v\xed\xf0F\x84-\x9b'S\xb2\xe7\x80"
const testnetAllocData = ""
const rinkebyAllocData = ""
