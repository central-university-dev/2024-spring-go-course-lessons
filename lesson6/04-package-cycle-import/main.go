package main

import (
	"04-package-cycle-import-issue/boo"
	"04-package-cycle-import-issue/foo"
)

func main() {

	b := boo.New()
	f := foo.New()

	f.Hello()
	b.Hello()
}
