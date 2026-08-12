//go:build !unix

package watchmanager

import "os"

// reRaiseSignal is a no-op here: these platforms cannot re-deliver a termination
// signal to the current process (on Windows, os.Process.Signal rejects Interrupt).
// The watch loop returns instead, and tsgo exits normally.
func reRaiseSignal(sig os.Signal) {
}
