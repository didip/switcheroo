[![GoDoc](https://godoc.org/github.com/didip/switcheroo?status.svg)](http://godoc.org/github.com/didip/switcheroo)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/didip/switcheroo/master/LICENSE)

## Switcheroo

Have you ever want to use an HTTP multiplexer for other things?

If you have, then Switcheroo is for you.


## 1 Minute Example

```go
import (
	"context"
	"strconv"
	"log"
)

ctx := context.WithValue(context.Background(), "total", 1)

r := New(ctx)

r.Add("/add/{number}", func(ctx context.Context, params map[string]string, others ...interface{}) {
	total := ctx.Value("total").(int)
	number, _ := strconv.Atoi(params["number"])
	total += number

	if total != 10 {
		log.Fatalf("total should have been 10. Got: %v", total)
	}
})

r.Run("/add/9")
```

## What is the string pattern used?

* `{name}` Named capture group delimited by curly braces.

* `*` to denote non-greedy globbing.



## Why do I need this?

Perhaps you want to multiplex based on GRPC or JSON-RPC methods?


## My other Go libraries

* [Tollbooth](https://github.com/didip/tollbooth): A generic middleware to rate-limit HTTP requests.

* [LaborUnion](https://github.com/didip/laborunion): A dynamic worker pool library.

* [Stopwatch](https://github.com/didip/stopwatch): A small library to measure latency of things. Useful if you want to report latency data to Graphite.
