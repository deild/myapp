package commands

import (
	"testing"
)

func TestFoo(t *testing.T) {
	err := Foo()
	if err != nil {
		t.Fatalf("Expected err nil but got %s", err)
	}
}
