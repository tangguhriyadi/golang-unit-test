package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HellowWorld("Tangguh")

	t.Run("Riyadi", func(t *testing.T) {
		result := HellowWorld("Riyadi")
		if result != "Hello Riyadi" {
			t.Error("NAHLO")
		}
	})

	assert.Equal(t, "Hello Tangguh", result)

	fmt.Println("Executed")
}

func TestMain(m *testing.M) {
	fmt.Println("sebelum")

	m.Run()

	fmt.Println("Setelah")
}
