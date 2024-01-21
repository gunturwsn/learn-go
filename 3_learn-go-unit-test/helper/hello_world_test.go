package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

/**
*	1. cd to package test
*	2. command 'go test' -> running test by package
*	3. command 'go test -v' -> running test each file detail test
*	4. command 'go test -v -run <function_name>' -> running test each function detail test
* 	(when running go test -v -run <function_name>, if there are function with same prefix
*	e.g in package there are two function with name TestHelloWorld & TestHelloWorldBudi
*	, so the two test function will be run).
*	5. command 'go test -v ./...' -> running test from previous package to package test
**/

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// error
		t.Error("Result must be 'Hello Budi'")
	}
	fmt.Println("TestHelloWorldBudi Done")
}

func TestHelloWorldHari(t *testing.T) {
	result := HelloWorld("Hari")
	if result != "Hello Hari" {
		// error
		t.Fatal("Result must be 'Hello Hari'")
	}
	fmt.Println("TestHelloWorldHari Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	// assert will call Fail()
	result := HelloWorld("Budi")
	assert.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
	fmt.Println("TestHelloWorld with Assertion Done")
}

func TestHelloWorldRequire(t *testing.T) {
	// require will call FailNow()
	result := HelloWorld("Budi")
	require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on Linux OS")
	}

	result := HelloWorld("Budi")
	require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
}
