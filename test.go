package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// Define a byte slice
	data := []byte("AB")

	// Encode the byte slice to a hexadecimal string
	encoded := hex.EncodeToString(data)

	// Print the encoded hexadecimal string
	fmt.Println(encoded) // Output: 48656c6c6f2c20476f21
}
