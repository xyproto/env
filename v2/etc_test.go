package env

import (
	"testing"
)

func TestFindJavaInEtcEnvironment(_ *testing.T) {
	_, _ = EtcEnvironment("JAVA_HOME")
}

func TestFindKotlinInEtcEnvironment(_ *testing.T) {
	_, _ = EtcEnvironment("KOTLIN_HOME")
}
