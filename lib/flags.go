package lib

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func Flags(version string) {
	helpPtr := flag.Bool("help", false, "Show usage information")
	encodePtr := flag.String("encode", "", "Encode a string to Base64")
	decodePtr := flag.String("decode", "", "Decode a Base64 to string")
	flag.Parse()
	// Handle help flag
	if *helpPtr {
		fmt.Fprintf(os.Stderr, "Usage: %s v%s [-encode input-string] [-decode base64-string]\n", os.Args[0], version)
		fmt.Println("cli64 is for encode string to get base64 as an output and to decode base64")
		fmt.Println("made by nbebaw")
		fmt.Fprintf(os.Stderr, "\nFlags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Handle encode and decode flags
	if *encodePtr != "" {
		// Encode input string to Base64
		encoded := base64.StdEncoding.EncodeToString([]byte(*encodePtr))
		fmt.Println(encoded)
	} else if *decodePtr != "" {
		// Decode input Base64 string
		decoded, err := base64.StdEncoding.DecodeString(*decodePtr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error decoding Base64 string:", err)
			os.Exit(1)
		}
		fmt.Println(string(decoded))
	} else {
		// No flags specified
		fmt.Fprintf(os.Stderr, "Error: no flags specified\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [-encode input-string] [-decode base64-string]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", os.Args[0])
		os.Exit(1)
	}
}
