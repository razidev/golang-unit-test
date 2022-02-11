package helper

import (
	"fmt"
	"testing"
)

func TestHeloWolrd1(t *testing.T) {
	result := HelloWorld("Razi")

	if result != "Hello Razi" {
		// t.Fail()
		t.Error("Error, result should be 'Hello Razi'")
	}
	fmt.Println("TestHeloWolrd1 done")
}
