//go:build wasip1 && wasm
// +build wasip1,wasm

package isatty

func IsTerminal(ds uintptr) bool       { return false }
func IsCygwinTerminal(fd uintptr) bool { return false }
