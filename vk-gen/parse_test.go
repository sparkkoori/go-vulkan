package main

import "testing"

func TestParseTypeString(t *testing.T) {
	def := parseTypeString("void const *")
	deepPrint(def, 0)

	println("\n")

	def = parseTypeString("int (*[5])(void)")
	deepPrint(def, 0)
}
