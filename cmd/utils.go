package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func fixFilePath(path string) string {
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(home, path[2:])
	}
	return path
}

func reportError(err string, exit bool) {
	if err == "" {
		fmt.Fprintf(os.Stderr, "Unexpected error occured\n")
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", strings.ToUpper(string(err[0]))+err[1:])
	}
	if exit {
		os.Exit(1)
	}
}
