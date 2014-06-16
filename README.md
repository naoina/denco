# Denco [![Build Status](https://travis-ci.org/naoina/denco.png?branch=master)](https://travis-ci.org/naoina/denco)

The fast and flexible URL router for [Go](http://golang.org).

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
		{"/", &route{"root"}},
		{"/user/:id", &route{"user"}},
		{"/user/:name/:id", &route{"username"}},
		{"/static/*filepath", &route{"static"}},
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

### URL patterns

Denco's route matching strategy is "most nearly matching".

When routes `/:name` and `/alice` have been built, URI `/alice` matches the route `/alice`, not `/:name`.
Because URI `/alice` is more match with the route `/alice` than `/:name`.

For more example, when routes below have been built:

```
/user/alice
/user/:name
/user/:name/:id
/user/alice/:id
/user/:id/bob
```

Routes matching are:

```
/user/alice      => "/user/alice" (no match with "/user/:name")
/user/bob        => "/user/:name"
/user/naoina/1   => "/user/:name/1"
/user/alice/1    => "/user/alice/:id" (no match with "/user/:name/:id")
/user/1/bob      => "/user/:id/bob"   (no match with "/user/:name/:id")
/user/alice/bob  => "/user/alice/:id" (no match with "/user/:name/:id" and "/user/:id/bob")
```

## Limitation

Denco has some limitations below.

* Number of param records (such as `/:name`) must be less than 2^22
* Number of elements of internal slice must be less than 2^22

## Benchmarks

    cd $GOPATH/github.com/naoina/kocha-urlrouter
    go test -bench . -benchmem

## License

Denco is licensed under the MIT License.
