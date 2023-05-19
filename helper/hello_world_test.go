package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HellowWorld("Tangguh")

	assert.Equal(t, "Hello Tangguh", result)

	fmt.Println("Executed")
}
