package main

import (
	"runtime"

	cmd "github.com/haskito/cmd/haskitod/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
