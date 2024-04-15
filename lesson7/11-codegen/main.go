package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(cwd, "example.go")
	fset := token.NewFileSet() // a set of sources
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("error parse file: %s", err)
	}

	ast.Print(fset, node)
}
