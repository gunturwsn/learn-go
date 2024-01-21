package helper

import "testing"

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
		panic("Result is not Hello Budi")
	}
}

func TestHelloWorldHari(t *testing.T) {
	result := HelloWorld("Hari")
	if result != "Hello Hari" {
		// error
		panic("Result is not Hello Hari")
	}
}
