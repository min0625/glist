// Package glist implements a generic doubly linked list.
//
// It is designed to provide behavior and method semantics aligned with
// Go's standard library container/list package while enabling type-safe
// element values through generics.
//
// The zero value of List is ready to use.
package glist

import "github.com/min0625/glist/internal/list"

// Element is an element of a linked list.
type Element[T any] = list.ElementOf[T]

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List[T any] = list.ListOf[T]

// New returns an initialized list.
func New[T any]() *List[T] {
	return list.NewOf[T]()
}
