# go-x11-hash

[![Build Status](https://gitlab.com/samli88/go-x11-hash/badges/master/build.svg)](https://gitlab.com/samli88/go-x11-hash/pipelines)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/gitlab.com/samli88/go-x11-hash)

Implements the X11 hash and required functions in Go

## Usage

```
    package main

    import (
        "fmt"

        "gitlab.com/samli88/go-x11-hash"
    )

    func main() {
        hs, out := x11.New(), [32]byte{}
        hs.Hash([]byte("DASH"), out[:])
        fmt.Printf("%x \n", out[:])
    }
```

## Notes

The echo, simd and shavite implementations do not have 100% test coverage. A
full test on these requires the test to hash a blob of bytes that is several
gigabytes large.

## License

go-x11-hash is licensed under the [copyfree](http://copyfree.org) ISC License.

## Attribution/Credit

This work is based on an original work by Nitya Sattva originally found at
https://gitlab.com/nitya-sattva/go-x11. Only a few tweaks have been made to
rename some of the hash algorithms to align with those names used in Dash Core,
and fix permissions.
