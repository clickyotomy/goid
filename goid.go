// Package goid provides a way to retrieve
// the runtime ID of the calling goroutine.
package goid

import (
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
)

// implArch is a list of architectures that have
// assembly implementations of "goidFast".
var implArch = map[string]struct{}{
	"386":      {},
	"amd64":    {},
	"amd64p32": {},
	"arm":      {},
	"arm64":    {},
	"riscv64":  {},
}

// fastSupport is a helper to determine if the "fast"
// version of the function is supported or not.
func fastSupport() bool {
	_, ok := implArch[runtime.GOARCH]
	return ok
}

// goidFast returns the goroutine ID. This is defined externally,
// and has archirecture-specific implementations depending on the
// value of "runtime.GOARCH".
func goidFast() int64

// goidSlow is a fallback implementation for "goidFast". This
// function is called when the value of "runtime.GOARCH" is not
// present in "implArch" map.
func goidSlow() int64 {
	buf := debug.Stack()

	if len(buf) <= 0 {
		return int64(-1)
	}

	// Format: "goroutine ${N} [running]: ..."
	chunks := strings.Split(string(buf), " ")
	if len(chunks) < 2 {
		return int64(-1)
	}

	ret, err := strconv.Atoi(strings.TrimSpace(chunks[1]))
	if err != nil {
		return int64(-1)
	}

	return int64(ret)
}

// Get returns the runtime ID of the calling goroutine.
func Get() int64 {
	if fastSupport() {
		return goidFast()
	}

	return goidSlow()
}
