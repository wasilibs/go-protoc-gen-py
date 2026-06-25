package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-py",
		LibraryRepo: "bufbuild/protobuf-py",
		GoReleaser:  true,
	})
	boot.Main()
}
