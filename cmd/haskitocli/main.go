package main

import (
	"runtime"

	cmd "github.com/haskito/cmd/haskitocli/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
