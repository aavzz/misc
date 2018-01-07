/*
Package pipe exposes some higher-level primitives to
work with pipes
*/
package pipe

import (
	"os"
	"io"
)

// CheckStdin return true when Stdin is connected to a pipe,
// false otherwise
func CheckStdin() (bool, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	if fi.Mode() & os.ModeNamedPipe == 0 {
		return false, nil
	}
	return true, nil
}

// Read reads up to size bytes from stdin and
// returns its input as a string
func Read(size int) (string, error) {
	buf := make([]byte, size)
	num, err := os.Stdin.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(buf[:num]), nil
}
