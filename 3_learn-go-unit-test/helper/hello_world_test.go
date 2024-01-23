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

func TestMain(m *testing.M) {
	/**
	* TestMain only working in one package.
	* If you want to test outside the current package (e.g. helper),
	* you have to create another TestMain in the package you want to test.
	 */

	// before
	fmt.Println("Before unit test")

	m.Run()

	// after
	fmt.Println("After unit test")
}

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

func TestSubTest(t *testing.T) {
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
	})
	t.Run("Hari", func(t *testing.T) {
		result := HelloWorld("Hari")
		require.Equal(t, "Hai Hari", result, "Result must be 'Hai Hari'")
	})

	/**
	* if you only want to run one subtest, using 'go test -v -run=<TestFunction>/<SubTestName>'
	* example -> 'go test -v -run=TestSubTest/Hari'
	 */
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Budi)",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "HelloWorld(Hari)",
			request:  "Hari",
			expected: "Hai Hari",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	/* running benchmark command:
	* 'go test -v -bench=.' 									-> running all benchmark in package
	* 'go test -v -run=NotMathUnitTest -bench=.' 				-> running benchmark without unit test
	* 'go test -v -run=NotMathUnitTest -bench=<BenchmarkName>' 	-> running benchmark selected
	* 'go test -v -bench=. ./...' 								-> running benchmark from root package and running in all package
	 */
	for i := 0; i < b.N; i++ {
		HelloWorld("Budi")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	/** running sub benchmark command:
	* 'go test -v -run=NotMathUnitTest -bench=<BencmarkName>/<SubBenchmarkName>'
	* e.g: 'go test -v -run=NotMathUnitTest -bench=BenchmarkHelloWorldSub/Budi'
	 */

	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})

	b.Run("Hari", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hari")
		}
	})
}
