package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/logrusorgru/aurora"
)

var tempDir string

func init() {
	log.SetFlags(0)

	dir := path.Join(".", ".vk-gen-temp")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}
	tempDir = dir
}

func main() {
	flag.Parse()

	src := parse()
	fmt.Printf("Parsed %d nodes.\n", aurora.Green(len(src)))
	tgt := generate(src)
	fmt.Printf("Generated %d go nodes and %d c nodes.\n", aurora.Green(len(tgt.targetGo)), aurora.Green(len(tgt.targetC)))
	print(tgt)
}
