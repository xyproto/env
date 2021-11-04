package env

import (
	"os"
	"testing"
)

func TestContains(t *testing.T) {
	os.Setenv("ASDF", "123asdf123")
	if !Contains("ASDF", "asdf") {
		t.Fail()
	}
}

func TestHas(t *testing.T) {
	if Has("ASDFASDFASDF") {
		t.Fail()
	}
	os.Setenv("ASDFASDFASDF", "1")
	if !Has("ASDFASDFASDF") {
		t.Fail()
	}
}

func TestBool(t *testing.T) {
	if Bool("QWERTYQWERTY") {
		t.Fail()
	}
	os.Setenv("QWERTYQWERTY", "yes")
	if !Bool("QWERTYQWERTY") {
		t.Fail()
	}
}

func TestExpandUser(t *testing.T) {
	if ExpandUser("~/test") == "/tmp" {
		t.Fail()
	}
}
