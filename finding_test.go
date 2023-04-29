package greetings

import (
	"testing"
)

func TestFindUniq(t *testing.T) {
	arr := []float32{1.0, 2.0, 2.0}
	result, err := FindUniq(arr)
	if result != 1.0 || err != nil {
		t.Fatalf(`result is %v, expected 1.0`, result)
	}
}
