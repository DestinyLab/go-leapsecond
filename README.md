# go-leapsecond

go-leapsecond is the tool of leap seconds.

[![GoDoc](https://godoc.org/github.com/DestinyLab/go-leapsecond?status.svg)](https://godoc.org/github.com/DestinyLab/go-leapsecond) [![Go Report Card](https://goreportcard.com/badge/github.com/DestinyLab/go-leapsecond)](https://goreportcard.com/report/github.com/DestinyLab/go-leapsecond) [![Build Status](https://travis-ci.org/DestinyLab/go-leapsecond.svg?branch=master)](https://travis-ci.org/DestinyLab/go-leapsecond) [![Coverage Status](https://coveralls.io/repos/github/DestinyLab/go-leapsecond/badge.svg?branch=master)](https://coveralls.io/github/DestinyLab/go-leapsecond?branch=master)

## Installation

```
go get -u github.com/DestinyLab/go-leapsecond
```

## Usage

```go
package main

import (
  "fmt"
  "time"

  "github.com/DestinyLab/go-leapsecond"
)

func main() {
  t1 := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
  fmt.Printf("%v", leapsecond.Have(t1))
  // Output: 27s

  t2 := time.Date(2010, 12, 1, 0, 0, 0, 0, time.UTC)
  fmt.Printf("%v", leapsecond.Diff(t1, t2))
  // Output: 3s
}
```

## More Info

- [Leap second](https://en.wikipedia.org/wiki/Leap_second)
