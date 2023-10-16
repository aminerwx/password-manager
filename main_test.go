package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	out, in := createUI()
	if out.Text != "Hello World!" {
		t.Error("Incorrect initial greeting")
	}

	test.Type(in, "Amine")

	if out.Text != "Hello Amine!" {
		t.Error("Incorrect greeting.")
	}
}
