package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HellowWorld("Tangguh")

	if result != "Hello Tangguh" {
		//unit test failed
		t.Fail()
	}
	fmt.Println("TestHelloWorld Done")
}
