//go:build !windows

package cli

import (
	"os"
	"syscall"
)

// requestStop asks proc to shut down on its own. On Unix that is SIGTERM.
// The caller still enforces the grace period and escalates to os.Kill.
func requestStop(proc *os.Process) error {
	return signalProcess(proc, syscall.SIGTERM)
}
