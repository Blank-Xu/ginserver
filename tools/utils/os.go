package utils

import "os"

func IsDir(name string) bool {
	f, err := os.Stat(name)
	if err != nil {
		return false
	}
	return f.IsDir()
}

func IsFile(name string) bool {
	return !IsDir(name)
}

func IsFileExit(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
