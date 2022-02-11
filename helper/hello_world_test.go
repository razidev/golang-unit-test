package helper

import (
	"fmt"
	"testing"
)

/* do not use panic, use t.Fail(), t.FailNow(),
t.Error(message string), t.Fatal(message string),
t.FatalNow(message string) instead */

func TestHeloWolrd1(t *testing.T) {
	result := HelloWorld("Razi")

	if result != "Hello Razi" {
		t.Error("Error, result should be 'Hello Razi'")
	}
	fmt.Println("TestHeloWolrd1 done")
}

func TestHeloWolrd2(t *testing.T) {
	result := HelloWorld("Hulk")

	if result != "HelloHulk" {
		// t.FailNow()
		t.Fatal("Error, result should be 'HelloHulk'")
	}
	fmt.Println("TestHeloWolrd2 done")
}
