package main

import "testing"

func TestParseTypeString(t *testing.T) {
	def := parseTypeString("void const *")
	deepPrint(def, 0)
}
