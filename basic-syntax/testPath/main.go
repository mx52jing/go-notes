package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if execPath, err := os.Executable(); err == nil {
		fmt.Printf(filepath.Join(execPath, "a", "b/c"))
	}
}