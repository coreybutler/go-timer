# go-timer

[![Version](https://img.shields.io/github/tag/coreybutler/go-timer.svg)](https://github.com/coreybutler/go-timer)
[![GoDoc](https://godoc.org/github.com/coreybutler/go-timer?status.svg)](https://godoc.org/github.com/coreybutler/go-timer)
[![Build Status](https://travis-ci.org/coreybutler/go-timer.svg?branch=master)](https://travis-ci.org/coreybutler/go-timer)

Use timers and intervals like you ould in JavaScript, but in Go.

**Install** `go get github.com/coreybutler/go-timer`

### Example Usage

```go
package main

import (
  "fmt"
  "github.com/coreybutler/go-timer"
  "time"
)

func main () {
  t := timer.SetInterval(func (args ...interface{}) {
    fmt.Println(args)
  }, 1000, "hello world")

  time.Sleep(5 * time.Second)
}
```

After running for 5 seconds, the script above would output:

```sh
hello world
hello world
hello world
hello world
```

---

#### Why?

Most of my day-to-day code is written in JavaScript. My primary motivator was creating a lightweight module that provided a similar coding experience to JavaScript's `setTimeout` and `setInterval`. Of course, Go is different, but this package provides an API that is more familiar to JavaScript developers working with Go.

Additionally, writing timers with Go is a tiny bit pedantic. Wrapping Go's ticker functionality makes the code more readable and easier to understand.
