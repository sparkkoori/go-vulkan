package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParseTypeString(t *testing.T) {
	if err := os.Chdir("testdata"); err != nil {
		fmt.Println("Chdir error:", err)
	}

	def := parseTypeString("void const *")
	deepPrint(def, 0)

	println("\n")

	def = parseTypeString("int (*[5])(void)")
	deepPrint(def, 0)
}
