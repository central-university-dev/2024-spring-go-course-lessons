package main

import (
	"03-package-cycle-import-issue/boo"
	"03-package-cycle-import-issue/foo"
)

func main() {

	b := boo.New()
	f := foo.New()

	f.Hello()
	b.Hello()
}
