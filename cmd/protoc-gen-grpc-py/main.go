package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-py/internal/runner"
)

func main() {
	os.Exit(runner.Run("protoc-gen-grpc-py", os.Args[1:], os.Stdin, os.Stdout, os.Stderr, "."))
}
