package main

import (
	"fmt"

	"gitlab.tcsbank.ru/a.klakovskiy/twosum"
)

// gitlab.tcsbank.ru/a.klakovskiy/twosum

func main() {
	fmt.Println("calc using module twosum:", twosum.TwoSum(1, 2))
	fmt.Println("calc using module twosum:", twosum.FourSum(1, 2, 3, 4))
}
