package main

import (
	"os"
	"politicker/politicker"
)

func main() {
	// 1. Hit members API
	// 2. Parse json into Go structs
	// 3. Store parsed json into postgres db

	os.Exit(politicker.CLI(os.Args[1:]))
}
