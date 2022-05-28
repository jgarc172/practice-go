package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3")
	expected := 3

	got := count(b)
	if got != expected {
		t.Errorf("Expected %d, got %d \n", expected, got)
	}
}
