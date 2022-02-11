package helper

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
		t.Fatal("Error, result should be 'HelloHulk'")
	}
	fmt.Println("TestHeloWolrd2 done")
}

//Using Assert == t.Fail()
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Razi")
	assert.Equal(t, "Hello Razi", result, "Error, result should be 'Hello Razi'")
	fmt.Println("TestHelloWorldAssert is done")
}

//Using Require == t.FailNow()
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Razi")
	require.Equal(t, "Hello Razi", result, "Error, result should be 'Hello Razi'")
	fmt.Println("TestHelloWorldRequire is not execute here if error")
}
