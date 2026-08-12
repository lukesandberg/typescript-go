package main

import (
	"context"
	"os"

	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/execute"
)

func main() {
	os.Exit(runMain())
}

func runMain() int {
	core.ApplyDebugStackLimit()
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "--lsp":
			return runLSP(args[1:])
		case "--api":
			return runAPI(args[1:])
		}
	}
	// No SIGINT/SIGTERM handler: like the JS tsc, the default disposition terminates
	// the process, so an interrupt stops the compile at once and yields the
	// conventional exit code (130/143) instead of running to completion and reporting
	// success. Watch mode blocks forever, so WatchManager.RunLoop observes the signal
	// itself and re-raises it; see the comment there.
	result := execute.CommandLine(context.Background(), newSystem(), args, nil)
	return int(result.Status)
}
