package glist_test

import (
	"fmt"

	"github.com/min0625/glist"
)

func ExampleNew() {
	l := glist.New[string]()
	l.PushBack("a")
	l.PushBack("b")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// a
	// b
}

func ExampleList_PushFront() {
	l := glist.New[int]()
	l.PushFront(2)
	l.PushFront(1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
}
