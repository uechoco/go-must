[![Go Reference](https://pkg.go.dev/badge/github.com/uechoco/go-must.svg)](https://pkg.go.dev/github.com/uechoco/go-must)
[![Go Report Card](https://goreportcard.com/badge/github.com/uechoco/go-must)](https://goreportcard.com/report/github.com/uechoco/go-must)

# go-must
general utilities for "Must" pattern. It simplifies safe initialization, handling errors in test, and so on.

## Installation

This package uses generics, so you need to use Go 1.18 or later.

```bash
go get github.com/uechoco/go-must
```

## Example Usage

```go
import (
	"time"
	"github.com/uechoco/go-must"
)

func TestExample(t *testing.T) {
	tm := must.Get2(time.Parse(time.RFC3339, "2020-01-01T00:00:00Z"))
	fmt.Println(tm)
}
```
