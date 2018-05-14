package commands

import (
	"fmt"
	"time"
)

// Foo is my foo function
func Foo() error {
	fmt.Println("Wait 2 seconds ...")
	time.Sleep(2 * time.Second)
	return nil
}
