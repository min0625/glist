// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.google file.

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *List):
//
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
//
//nolint:revive
package list

type Element = ElementOf[any]

type List = ListOf[any]

func New() *List {
	return NewOf[any]()
}
