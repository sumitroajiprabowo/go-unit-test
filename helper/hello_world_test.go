package helper

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		expected string
		actual   string
	}{
		{"World", "Hello World", HelloWorld("World")},
		{"Dunia", "Hello Dunia", HelloWorld("Dunia")},
		{"Keras", "Hello Keras", HelloWorld("Keras")},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.actual = HelloWorld(bm.name)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("World", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("World")
		}
	})
	b.Run("Dunia", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Dunia")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("World")
	}
}

func BenchmarkHelloDunia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dunia Keras")
	}
}

func TestSubTest(t *testing.T) {
	t.Run("World", func(t *testing.T) {
		actual := HelloWorld("World")
		expected := "Hello World"
		assert.Equal(t, expected, actual, "Result must be equal")
	})
	t.Run("Dunia", func(t *testing.T) {
		actual := HelloWorld("Dunia")
		expected := "Hello Dunia"
		require.Equal(t, expected, actual, "Result must be equal")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Running tests...")
	result := m.Run()
	fmt.Println("Done.")
	os.Exit(result)
}

// Create a new test case with the name TestSkip with a skip statement
func TestSkip(t *testing.T) { // t is a pointer to a testing.T
	if runtime.GOOS == "linux" { // if runtime.GOOS is equal to linux
		t.Skip("Skipping test on linux") // t.Skip is a function that takes a format string and a variable number of arguments and marks the test as skipped.
	}
	actual := HelloWorld("World")                              // actual is a string
	expected := "Hello World"                                  // expected is a string
	require.Equal(t, expected, actual, "Result must be equal") // require.Equal is a function that takes two values and returns true if they are equal.
	fmt.Println(t.Name() + "with Require has done")            // fmt.Printf is a function that takes a format string and a variable number of arguments and returns a string.
}

// Create a new test case with the name TestHelloWorld
func TestHelloWorld(t *testing.T) { // t is a pointer to a testing.T
	expected := "Hello World"     // expected is a string
	actual := HelloWorld("World") // actual is a string
	if actual != expected {       // if actual is not equal to expected
		t.Errorf("HelloWorld(\"World\") = %s, expected %s", actual, expected) // t.Errorf is a function that takes a format string and a variable number of arguments and prints the formatted string to the testing.T.
	}
	fmt.Printf("Successfully ran test: %s\n", t.Name()) // fmt.Printf is a function that takes a format string and a variable number of arguments and returns a string.
}

// Create a new test case with the name TestHelloWorldFail
func TestFailHelloWorld(t *testing.T) { // t is a pointer to a testing.T
	expected := "Hello World"     // expected is a string
	actual := HelloWorld("World") // actual is a string
	if actual != expected {       // if actual is not equal to expected
		t.Errorf("HelloWorld(\"World\") = %s, expected %s", actual, expected) // t.Errorf is a function that takes a format string and a variable number of arguments and prints the formatted string to the testing.T.
	}
	fmt.Println("Successfully ran test:", t.Name()) // fmt.Printf is a function that takes a format string and a variable number of arguments and returns a string.
}

// Create a new test case with the name TestHelloWorldRequire with a require statement
func TestHelloWorldRequire(t *testing.T) { // t is a pointer to a testing.T
	actual := HelloWorld("World")                              // actual is a string
	expected := "Hello World"                                  // expected is a string
	require.Equal(t, expected, actual, "Result must be equal") // require.Equal is a function that takes two values and returns true if they are equal.
	fmt.Println(t.Name() + "with Require has done")            // fmt.Printf is a function that takes a format string and a variable number of arguments and returns a string.
}

// Create a new test case with the name TestHelloWorldAssert with an assert statement
func TestHelloWorldAssertion(t *testing.T) { // t is a pointer to a testing.T
	actual := HelloWorld("World")                             // actual is a string
	expected := "Hello World"                                 // expected is a string
	assert.Equal(t, expected, actual, "Result must be equal") // assert.Equal is a function that takes two values and returns true if they are equal.
	fmt.Println("Successfully ran test:", t.Name())           // fmt.Printf is a function that takes a format string and a variable number of arguments and returns a string.
}

// Example using TableTest
func TestHelloWorldTableTest(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		actual   string
	}{
		{"World", "Hello World", HelloWorld("World")},
		{"Dunia", "Hello Dunia", HelloWorld("Dunia")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.actual)
		})
	}
}

// Another Example Using Table Test Again
func TestHelloWOrldTableTestAgain(t *testing.T) {
	var tests = []struct {
		name     string
		request  string
		expected string
	}{
		{name: "World", request: "World", expected: "Hello World"},
		{name: "Dunia", request: "Dunia", expected: "Hello Dunia"},
		{name: "Indonesia", request: "Indonesia", expected: "Hello Indonesia"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
