package ex6_test

import (
	"ex6"
	"testing"
)

func TestIsNumBig(t *testing.T) {
	if !ex6.IsNumBig(101) {
		t.Error("101 should be big")
	}
}
