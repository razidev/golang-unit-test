package helper

import (
	"fmt"
	"runtime"

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

	if result != "Hello Hulk" {
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

// skip the process after skip function called, still success
func TestSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Can not run except on Mac OS")
	}

	result := HelloWorld("Razi")
	require.Equal(t, "Hello Razi", result, "Error, result should be 'Hello Razi'")
	fmt.Println("TestSkip is done")
}

//this unit test is to run all unit test
func TestMain(m *testing.M) {
	fmt.Println(">>> Before Unit Test <<<")
	fmt.Println()

	m.Run()

	fmt.Println()
	fmt.Println(">>> After Unit Test <<<")
}
