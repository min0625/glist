# glist
[![Go Reference](https://pkg.go.dev/badge/github.com/min0625/glist.svg)](https://pkg.go.dev/github.com/min0625/glist)

`glist` is a generic doubly linked list for Go with an API and behavior aligned
with `container/list`.

It is intended for users who want the same list operations as the standard
library list while getting compile-time type safety.

## Install

```bash
go get github.com/min0625/glist
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/min0625/glist"
)

func main() {
	l := glist.New[int]()
	l.PushBack(10)
	l.PushBack(20)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
```

## API Compatibility Goal

`glist` mirrors the `container/list` API shape where possible:

- `List[T]` and `Element[T]` correspond to `list.List` and `list.Element`
- Method names and semantics are the same as `container/list`
- Zero-value list behavior is preserved

## License

This project is licensed under Apache License 2.0. See `LICENSE`.

Parts of the implementation are derived from Go's standard library
`container/list` and are redistributed under the BSD-3-Clause terms from
the Go project. See `LICENSE.google` and `NOTICE` for attribution details.
