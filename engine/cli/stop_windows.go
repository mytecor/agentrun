//go:build windows

package cli

import "os"

// requestStop is a no-op on Windows, where os.Process.Signal supports only
// Kill and there is no way to ask a process to exit gracefully. Callers close
// stdin before calling this, which is what a well-behaved agent reacts to, and
// escalate to os.Kill once the grace period expires.
func requestStop(*os.Process) error {
	return nil
}
