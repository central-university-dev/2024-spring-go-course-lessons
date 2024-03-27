package boo

import (
	"fmt"

	"03-package-cycle-import-issue/foo"
)

type Boo struct {
	eventOneHandler *foo.Foo
}

func New() *Boo {
	return &Boo{}
}

func (b *Boo) Hello() {
	fmt.Println("Boo say hello")
}
