// Copyright 2019 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Package cloudpb is a subset of types and functions, copied from cloud.google.com/go/datastore.
//
// They are copied here to provide compatibility to decode keys generated by the cloud.google.com/go/datastore package.
package cloudkey

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/sgtsquiggs/protobuf/proto"
	cloudpb "github.com/sgtsquiggs/appengine/datastore/internal/cloudpb"
)

/////////////////////////////////////////////////////////////////////
// Code below is copied from https://github.com/googleapis/google-cloud-go/blob/master/datastore/datastore.go
/////////////////////////////////////////////////////////////////////

var (
	// ErrInvalidKey is returned when an invalid key is presented.
	ErrInvalidKey = errors.New("datastore: invalid key")
)

/////////////////////////////////////////////////////////////////////
// Code below is copied from https://github.com/googleapis/google-cloud-go/blob/master/datastore/key.go
/////////////////////////////////////////////////////////////////////

// Key represents the datastore key for a stored entity.
type Key struct {
	// Kind cannot be empty.
	Kind string
	// Either ID or Name must be zero for the Key to be valid.
	// If both are zero, the Key is incomplete.
	ID   int64
	Name string
	// Parent must either be a complete Key or nil.
	Parent *Key

	// Namespace provides the ability to partition your data for multiple
	// tenants. In most cases, it is not necessary to specify a namespace.
	// See docs on datastore multitenancy for details:
	// https://cloud.google.com/datastore/docs/concepts/multitenancy
	Namespace string
}

// DecodeKey decodes a key from the opaque representation returned by Encode.
func DecodeKey(encoded string) (*Key, error) {
	// Re-add padding.
	if m := len(encoded) % 4; m != 0 {
		encoded += strings.Repeat("=", 4-m)
	}

	b, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	pKey := new(cloudpb.Key)
	if err := proto.Unmarshal(b, pKey); err != nil {
		return nil, err
	}
	return protoToKey(pKey)
}

// valid returns whether the key is valid.
func (k *Key) valid() bool {
	if k == nil {
		return false
	}
	for ; k != nil; k = k.Parent {
		if k.Kind == "" {
			return false
		}
		if k.Name != "" && k.ID != 0 {
			return false
		}
		if k.Parent != nil {
			if k.Parent.Incomplete() {
				return false
			}
			if k.Parent.Namespace != k.Namespace {
				return false
			}
		}
	}
	return true
}

// Incomplete reports whether the key does not refer to a stored entity.
func (k *Key) Incomplete() bool {
	return k.Name == "" && k.ID == 0
}

// protoToKey decodes a protocol buffer representation of a key into an
// equivalent *Key object. If the key is invalid, protoToKey will return the
// invalid key along with ErrInvalidKey.
func protoToKey(p *cloudpb.Key) (*Key, error) {
	var key *Key
	var namespace string
	if partition := p.PartitionId; partition != nil {
		namespace = partition.NamespaceId
	}
	for _, el := range p.Path {
		key = &Key{
			Namespace: namespace,
			Kind:      el.Kind,
			ID:        el.GetId(),
			Name:      el.GetName(),
			Parent:    key,
		}
	}
	if !key.valid() { // Also detects key == nil.
		return key, ErrInvalidKey
	}
	return key, nil
}
