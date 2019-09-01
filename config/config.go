// Package config : Configuration file and directory
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package config

import (
	"os"
	"path/filepath"
	"strings"
)

func Path(appName, fileName string) string {
	if len(fileName) == 0 || includeSlash(fileName) {
		return ""
	}
	dir := Dir(appName)
	if len(dir) == 0 {
		return ""
	}
	return filepath.Join(dir, fileName)
}

//Dir returns Configuration directory
func Dir(appName string) string {
	if includeSlash(appName) {
		return ""
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = ""
	}
	if len(dir) == 0 {
		dir, err = os.UserHomeDir()
		if err != nil {
			dir = ""
		}
	}
	if len(dir) == 0 {
		return ""
	}
	return filepath.Join(dir, appName)
}

func includeSlash(path string) bool {
	if len(path) == 0 {
		return false
	}
	return strings.Contains(filepath.ToSlash(path), "/")
}
