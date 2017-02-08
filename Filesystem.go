package main

import (
	"fmt"
	"os"
	"strings"
)

func MakeAllDirs(path string) error {
	return os.MkdirAll(AbsPath(path), 0644)
}

func AbsPath(path string) string {
	return fmt.Sprint("%s/%s", bin_path, strings.Trim(path, "/"))
}
