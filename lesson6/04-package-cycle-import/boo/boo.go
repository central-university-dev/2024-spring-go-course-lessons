package boo

import (
	"fmt"
)

type Helloer interface {
	Hello()
}

type Boo struct {
	eventOneHandler Helloer
}

func New() *Boo {
	return &Boo{}
}

func (b *Boo) Hello() {
	fmt.Println("Boo say hello")
}
