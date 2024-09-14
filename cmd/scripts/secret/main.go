package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func generateSecret() string {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

func main() {
	fmt.Println(generateSecret())
}
