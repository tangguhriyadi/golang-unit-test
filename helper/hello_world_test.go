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
