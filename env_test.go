package env

import (
	"os"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	Unload()
	Unset("ASDF")
	os.Setenv("ASDF", "123asdf123")
	if !Contains("ASDF", "asdf") {
		t.Fail()
	}
	Unset("ASDF")
}

func TestHas(t *testing.T) {
	Unload()
	Unset("ASDFASDFASDF")
	if Has("ASDFASDFASDF") {
		t.Fail()
	}
	os.Setenv("ASDFASDFASDF", "1")
	if !Has("ASDFASDFASDF") {
		t.Fail()
	}
	Unset("ASDFASDFASDF")
}

func TestBool(t *testing.T) {
	Unload()
	Unset("QWERTYQWERTY")
	if Bool("QWERTYQWERTY") {
		t.Fail()
	}
	os.Setenv("QWERTYQWERTY", "yes")
	if !Bool("QWERTYQWERTY") {
		t.Fail()
	}
	Unset("QWERTYQWERTY")
}

func TestExpandUser(t *testing.T) {
	Unload()
	if ExpandUser("~/test") == "/tmp" {
		t.Fail()
	}
}

func TestDir(t *testing.T) {
	Unload()
	Unset("P")
	if x := Dir("P", "~/ost"); x == "" || x == "/tmp" || strings.HasPrefix(x, "~") {
		t.Fail()
	}
	os.Setenv("P", "~/test")
	if x := Dir("P", "~/ost"); x == "" || x == "/tmp" || strings.HasPrefix(x, "~") {
		t.Fail()
	}
	Unset("P")
}

func TestPath(t *testing.T) {
	Unload()
	os.Setenv("PATH", "/usr/bin:/bin:/usr/local/bin")
	if len(Path()) != 3 {
		t.Fail()
	}
	Unset("PATH")
}

func TestHomeDir(t *testing.T) {
	if HomeDir() == "" {
		t.Fail()
	}
}

func TestFloat32(t *testing.T) {
	Unload()
	Set("F", "1.234")
	f := Float32("F", 0.0)
	if f != 1.234 {
		t.Fail()
	}
	Unset("F")
}
