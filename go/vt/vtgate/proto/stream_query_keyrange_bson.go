// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

// DO NOT EDIT.
// FILE GENERATED BY BSONGEN.

import (
	"bytes"

	"github.com/youtube/vitess/go/bson"
	"github.com/youtube/vitess/go/bytes2"
)

// MarshalBson bson-encodes StreamQueryKeyRange.
func (streamQueryKeyRange *StreamQueryKeyRange) MarshalBson(buf *bytes2.ChunkedWriter, key string) {
	bson.EncodeOptionalPrefix(buf, bson.Object, key)
	lenWriter := bson.NewLenWriter(buf)

	bson.EncodeString(buf, "Sql", streamQueryKeyRange.Sql)
	// map[string]interface{}
	{
		bson.EncodePrefix(buf, bson.Object, "BindVariables")
		lenWriter := bson.NewLenWriter(buf)
		for _k, _v1 := range streamQueryKeyRange.BindVariables {
			bson.EncodeInterface(buf, _k, _v1)
		}
		lenWriter.Close()
	}
	bson.EncodeString(buf, "Keyspace", streamQueryKeyRange.Keyspace)
	bson.EncodeString(buf, "KeyRange", streamQueryKeyRange.KeyRange)
	streamQueryKeyRange.TabletType.MarshalBson(buf, "TabletType")
	// *Session
	if streamQueryKeyRange.Session == nil {
		bson.EncodePrefix(buf, bson.Null, "Session")
	} else {
		(*streamQueryKeyRange.Session).MarshalBson(buf, "Session")
	}

	lenWriter.Close()
}

// UnmarshalBson bson-decodes into StreamQueryKeyRange.
func (streamQueryKeyRange *StreamQueryKeyRange) UnmarshalBson(buf *bytes.Buffer, kind byte) {
	switch kind {
	case bson.EOO, bson.Object:
		// valid
	case bson.Null:
		return
	default:
		panic(bson.NewBsonError("unexpected kind %v for StreamQueryKeyRange", kind))
	}
	bson.Next(buf, 4)

	for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
		switch bson.ReadCString(buf) {
		case "Sql":
			streamQueryKeyRange.Sql = bson.DecodeString(buf, kind)
		case "BindVariables":
			// map[string]interface{}
			if kind != bson.Null {
				if kind != bson.Object {
					panic(bson.NewBsonError("unexpected kind %v for streamQueryKeyRange.BindVariables", kind))
				}
				bson.Next(buf, 4)
				streamQueryKeyRange.BindVariables = make(map[string]interface{})
				for kind := bson.NextByte(buf); kind != bson.EOO; kind = bson.NextByte(buf) {
					_k := bson.ReadCString(buf)
					var _v1 interface{}
					_v1 = bson.DecodeInterface(buf, kind)
					streamQueryKeyRange.BindVariables[_k] = _v1
				}
			}
		case "Keyspace":
			streamQueryKeyRange.Keyspace = bson.DecodeString(buf, kind)
		case "KeyRange":
			streamQueryKeyRange.KeyRange = bson.DecodeString(buf, kind)
		case "TabletType":
			streamQueryKeyRange.TabletType.UnmarshalBson(buf, kind)
		case "Session":
			// *Session
			if kind != bson.Null {
				streamQueryKeyRange.Session = new(Session)
				(*streamQueryKeyRange.Session).UnmarshalBson(buf, kind)
			}
		default:
			bson.Skip(buf, kind)
		}
	}
}
