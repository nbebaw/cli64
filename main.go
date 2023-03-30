package main

import (
	"cli64/lib"
)

var Version = "No Version Provided"

func main() {
	// Define command-line flags
	lib.Flags(Version)
}
