package foo

import (
	"fmt"

	"04-package-cycle-import-issue/boo"
)

type Foo struct {
	eventTwoHandler *boo.Boo
}

func New() *Foo {
	return &Foo{}
}

func (f *Foo) Hello() {
	fmt.Println("Foo say hello")
}
