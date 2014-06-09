# Denco [![Build Status](https://travis-ci.org/naoina/denco.png?branch=master)](https://travis-ci.org/naoina/denco)

The fast URL router for [Go](http://golang.org).

Denco is based on Double-Array implementation of [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter), but does some optimizations for performance improvement.

## Features

* Fast (See [go-http-routing-benchmark](https://github.com/naoina/go-http-routing-benchmark))
* URL patterns (`/foo/:bar` and `/foo/*wildcard`)
* Minimum functions for maximum flexibility

Denco is **NOT** HTTP request multiplexer like Go's `http.ServeMux`, so doesn't provides Go's `http.Handler` interface.

## Installation

    go get -u github.com/naoina/denco

## Usage

```go
package main

import (
	"fmt"

	"github.com/naoina/denco"
)

type route struct {
	name string
}

func main() {
	router := denco.New()
	router.Build([]denco.Record{
		denco.NewRecord("/", &route{"root"}),
		denco.NewRecord("/user/:id", &route{"user"}),
		denco.NewRecord("/user/:name/:id", &route{"username"}),
		denco.NewRecord("/static/*filepath", &route{"static"}),
	})

	data, params, found := router.Lookup("/")
	// print `&main.route{name:"root"}, []denco.Param(nil), true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/user/hoge")
	// print `&main.route{name:"user"}, []denco.Param{denco.Param{Name:"id", Value:"hoge"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/user/hoge/7")
	// print `&main.route{name:"username"}, []denco.Param{denco.Param{Name:"name", Value:"hoge"}, denco.Param{Name:"id", Value:"7"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)

	data, params, found = router.Lookup("/static/path/to/file")
	// print `&main.route{name:"static"}, []denco.Param{denco.Param{Name:"filepath", Value:"path/to/file"}}, true`.
	fmt.Printf("%#v, %#v, %#v\n", data, params, found)
}
```

See [Godoc](http://godoc.org/github.com/naoina/denco) for more details.

## Benchmarks

    cd $GOPATH/github.com/naoina/kocha-urlrouter
    go test -bench . -benchmem

## License

Denco is licensed under the MIT License.
