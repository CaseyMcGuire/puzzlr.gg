package util

import "os"

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	// Return false if it's a directory, true otherwise
	return !info.IsDir()
}
