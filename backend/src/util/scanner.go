package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Returns a scanner with the offset set to the start of the input.
func NewScannerFromOffset(input io.Reader, start int) (*bufio.Scanner, error) {
	if start < 0 {
		return nil, fmt.Errorf("start offset cannot be negative")
	}

	scanner := bufio.NewScanner(input)
	for i := 0; i < start; i++ {
		scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return scanner, nil
}

// Returns a slice of logs from the input file starting from the given offset
func ScanFileFromOffset(f *os.File, offset int) (logs []string, err error) {
	scanner, err := NewScannerFromOffset(f, offset)
	if err != nil {
		return nil, err
	}
	logs = make([]string, 0)
	defer func() {
		// Recover from panic from scanner.Scan()
		if recovery := recover(); recovery != nil {
			err = fmt.Errorf("function panicked while scanning logs: %s", recovery)
			logs = nil
		}
	}()

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		logs = append(logs, scanner.Text())
	}
	return logs, nil
}
