package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-py",
		LibraryRepo: "bufbuild/protoc-gen-py",
		GoReleaser:  true,
	})
	boot.Main()
}
