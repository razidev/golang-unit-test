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

func TestSubTest(t *testing.T) {
	t.Run("Razi", func(t *testing.T) {
		result := HelloWorld("Razi")
		require.Equal(t, "Hello Razi", result, "Error, result should be 'Hello Razi'")
	})

	t.Run("Syahputro", func(t *testing.T) {
		result := HelloWorld("Syahputro")
		require.Equal(t, "Hello Syahputro", result, "Error, result should be 'Hello Syahputro'")
	})

	//if you want to run each subtest, e.g. go test -v -run=TestSubTest/Razi
}

func TestTable(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "Batman",
			request:  "Batman",
			expected: "Hello Batman",
		},
		{
			name:     "Superman",
			request:  "Superman",
			expected: "Hello Superman",
		},
		{
			name:     "Aquaman",
			request:  "Aquaman",
			expected: "Hello Aquaman",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Error, result should be"+test.expected)
		})
	}
}

//how to run benchmark
//go test -v -run=TestNotMatch -bench=Benchmarkxxx
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Razi")
	}
}

func BenchmarkHelloWorldSyahputro(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Syahputro")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Razi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Razi")
		}
	})

	b.Run("Syahputro", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Syahputro")
		}
	})
}
