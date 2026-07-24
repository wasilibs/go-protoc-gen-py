# go-protoc-gen-py

go-protoc-gen-py is a distribution of [protoc-gen-py][1], that can be built with Go. It does not actually reimplement any
functionality of protoc-gen-py in Go, instead packaging it with the WASI build of [Python][3], and
executing with the pure Go Wasm runtime [wazero][2]. This means that `go install` or `go run`
can be used to execute it, with no need to rely on separate package managers such as pnpm,
on any platform that Go supports.

## Installation

Precompiled binaries are available in the [releases](https://github.com/wasilibs/go-protoc-gen-py/releases).
Alternatively, install the plugin you want using `go install`.

```bash
$ go install github.com/wasilibs/go-protoc-gen-py/cmd/protoc-gen-py@latest
```

To avoid installation entirely, it can be convenient to use `go run`

```bash
$ go run github.com/wasilibs/go-protoc-gen-py/cmd/protoc-gen-py@latest .
```

[1]: https://github.com/bufbuild/protoc-gen-py
[2]: https://wazero.io/
[3]: https://github.com/brettcannon/cpython-wasi-build
