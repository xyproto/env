package env

import (
	"os"
	"strings"
	"testing"
)

func TestCachingContains(t *testing.T) {
	os.Setenv("ASDF", "123asdf123")
	Load()
	if !Contains("ASDF", "asdf") {
		t.Fail()
	}
	Unset("ASDF")
}

func TestCachingHas(t *testing.T) {
	Load()
	if Has("ASDFASDFASDF") {
		t.Fail()
	}
	os.Setenv("ASDFASDFASDF", "1")
	Load()
	if !Has("ASDFASDFASDF") {
		t.Fail()
	}
	Unset("ASDFASDFASDF")
}

func TestCachingBool(t *testing.T) {
	Load()
	if Bool("QWERTYQWERTY") {
		t.Fail()
	}
	os.Setenv("QWERTYQWERTY", "yes")
	Load()
	if !Bool("QWERTYQWERTY") {
		t.Fail()
	}
	Unset("QERTYQWERTY")
}

func TestCachingExpandUser(t *testing.T) {
	Load()
	if ExpandUser("~/test") == "/tmp" {
		t.Fail()
	}
}

func TestCachingDir(t *testing.T) {
	Load()
	if x := Dir("P", "~/ost"); x == "" || x == "/tmp" || strings.HasPrefix(x, "~") {
		t.Fail()
	}
	os.Setenv("P", "~/test")
	Load()
	if x := Dir("P", "~/ost"); x == "" || x == "/tmp" || strings.HasPrefix(x, "~") {
		t.Fail()
	}
	Unset("P")
}

func TestCachingPath(t *testing.T) {
	os.Setenv("PATH", "/usr/bin:/bin:/usr/local/bin")
	Load()
	if len(Path()) != 3 {
		t.Fail()
	}
	Unset("PATH")
}

func TestCachingHomeDir(t *testing.T) {
	Load()
	if HomeDir() == "" {
		t.Fail()
	}
}
