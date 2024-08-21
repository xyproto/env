package main

import (
	"fmt"

	"github.com/xyproto/env/v2"
)

func main() {
	envMap := env.Map()
	sortedKeys := env.Keys()

	// Output all set environment variables, sorted by the name
	for _, key := range sortedKeys {
		fmt.Println(key + " = " + envMap[key])
	}
}
