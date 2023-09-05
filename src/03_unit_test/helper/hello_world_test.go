package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Ari",
			request: "Ari",
		},
		{
			name:    "Wibowo",
			request: "Wibowo",
		},
	}

	for _, bencbenchmark := range benchmarks {
		b.Run(bencbenchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(bencbenchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Dava", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Dava")
		}
	})
	b.Run("Raihan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Raihan")
		}
	})
}

func BenchmarkSayFajar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Fajar")
	}
}

func BenchmarkSayAhmad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Ahmad")
	}
}

func TestMain(m *testing.M) {
	// before execute func
	fmt.Println("Before Unit Test")

	m.Run()

	// After execute func
	fmt.Println("After Unit Test")
}

// testing with table
func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fajar",
			request:  "Fajar",
			expected: "Hello Fajar",
		},
		{
			name:     "Ahmad",
			request:  "Ahmad",
			expected: "Hello Ahmad",
		},
		{
			name:     "Octhaviano",
			request:  "Octhaviano",
			expected: "Hello Octhaviano",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Fajar", func(t *testing.T) {
		result := SayHello("Fajar")

		// if fail will use failnow()
		require.Equal(t, "Hello Fajar", result, "Result must be 'Hello Fajar'")
	})
	t.Run("Ahmad", func(t *testing.T) {
		result := SayHello("Ahmad")

		// if fail will use failnow()
		require.Equal(t, "Hello Ahmad", result, "Result must be 'Hello Ahmad'")
	})
}

func TestSayHelloRequire(t *testing.T) {
	result := SayHello("Fajar")

	// if fail will use failnow()
	require.Equal(t, "Hello Fajar", result, "Result must be 'Hello Fajar'")
}

func TestSayHelloAssert(t *testing.T) {
	result := SayHello("Fajar")

	// if fail will use fail()
	assert.Equal(t, "Hello Fajar", result, "Result must be 'Hello Fajar'")
}

func TestSayHello(t *testing.T) {
	result := SayHello("Fajar")

	if result != "Hello Fajar" {
		// Error and will use fail()
		t.Error("Result must be 'Hello Fajar'")
	}
}
