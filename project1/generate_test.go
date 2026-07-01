package main

import "testing"

func TestGenerateSymbols(t *testing.T) {
	code, err := generateSymbols()
	if err != nil {
		t.Fatal(err)
	}
	if len(code) != 6 {
		t.Fatalf("ожидал 6 символов, получил %d: %q", len(code), code)
	}
	t.Logf("сгенерировано: %s", code)
}