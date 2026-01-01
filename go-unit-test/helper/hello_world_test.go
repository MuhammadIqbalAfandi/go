package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTableDdriven(b *testing.B) {
	tests := []struct {
		Name string
		Request string
	} {
		{
			Name: "Iqbal",
			Request: "Iqbal",
		},
		{
			Name: "Eko",
			Request: "Eko Kurniawan",
		},
	}

	for _, test := range tests {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.Request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Iqbal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Iqbal")
		}
	})

	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Iqbal")
	}
}

func BenchmarkHelloWorldEko(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func TestHelloWorldTableDriven(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {
		{
			name: "Iqbal",
			request: "Iqbal",
			expected: "Hello Iqbal",
		},
		{
			name: "Afandi",
			request: "Afandi",
			expected: "Hello Afandi",
		},
		{
			name: "Sanju",
			request: "Sanju",
			expected: "Hello Sanju",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestHelloWorldSubTest(t *testing.T) {
	t.Run("Iqbal", func(t *testing.T) {
		result := HelloWorld("Iqbal")
		require.Equal(t, "Hello Iqbal", result, "Result must be 'Hello Iqbal' but got '%s'", result)
	})

	t.Run("Iqbal", func(t *testing.T) {
		result := HelloWorld("Afandi")
		require.Equal(t, "Hello Afandi", result, "Result must be 'Hello Afandi' but got '%s'", result)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on Linux")
	}

	result := HelloWorld("Iqbal")
	if result != "Hello Iqbal" {
		t.Errorf("Result must be 'Hello Iqbal' but got '%s'", result)
	}

	fmt.Println("running in the end")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Iqbal")
	if result != "Hello Iqbal" {
		require.Equal(t, "Hello Iqbal", result, "Result must be 'Hello Iqbal'")
	}

	fmt.Println("running in the end")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Iqbal")
	if result != "Hello Iqbal" {
		assert.Equal(t, "Hello Iqbal", result, "Result must be 'Hello Iqbal'")
	}

	fmt.Println("running in the end")
}

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Iqbal")
	if result != "Hello Iqbal" {
		t.Fail()
	}

	fmt.Println("running in the end")
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Iqbal")
	if result != "Hello Iqbal" {
		t.FailNow()
	}

	fmt.Println("running in the end")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Iqbal")
	if result != "Hello Iqbal" {
		t.Error("Result not matched")
	}

	fmt.Println("running in the end")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Iqbal")
	if result != "Hello Iqbal" {
		t.Fatal("Result not matched")
	}

	fmt.Println("running in the end")
}