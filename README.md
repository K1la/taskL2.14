# or-channel

A small Go utility package that merges multiple **done channels** into one.  
The resulting channel is closed as soon as **any** of the input channels is closed.  

This pattern is useful for coordinating goroutines and handling cancellation in concurrent code.

---

## Installation

```bash
go get github.com/K1la/taskL2.14
````

---

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/K1la/taskL2.14/or"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or.Or(
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}

```

---

## Development

### Run example

```bash
make run
```

### Run tests

```bash
make test
```

---

## Project structure

```
.
├── or/
│   ├── or.go
│   └── or_test.go
├── main.go
├── go.mod 
└── README.md
```