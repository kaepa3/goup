package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	upper()
}

func upper() {
	f, err := os.Open("sample.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%x\n", h.Sum(nil))
}
