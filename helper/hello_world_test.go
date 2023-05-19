package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HellowWorld("Tangguh")

	t.Run("Riyadi", func(t *testing.T) {
		result := HellowWorld("Riyadi")
		require.Equal(t, "Hello Riyadi", result)
	})

	assert.Equal(t, "Hello Tangguh", result)

	fmt.Println("Executed")
}

func TestMain(m *testing.M) {
	fmt.Println("sebelum")

	m.Run()

	fmt.Println("Setelah")
}

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Muhammad)",
			request:  "Muhammad",
			expected: "Hello Muhammad",
		},
		{
			name:     "HelloWorld(Tangguh)",
			request:  "Tangguh",
			expected: "Hello Tangguh",
		},
		{
			name:     "HelloWorld(Riyadi)",
			request:  "Riyadi",
			expected: "Hello Riyadi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HellowWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
