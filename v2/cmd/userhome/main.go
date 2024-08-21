package main

import (
	"fmt"

	"github.com/xyproto/env/v2"
)

func main() {
	fmt.Printf("Current user:\t\t%s\n", env.CurrentUser())
	fmt.Printf("Home directory:\t\t%s\n", env.HomeDir())
	fmt.Printf("$SHELL:\t\t\t%s\n", env.Str("SHELL", "unknown"))
	fmt.Printf("$BROWSER:\t\t%s\n", env.Str("BROWSER", "unknown"))
	fmt.Printf("$EDITOR:\t\t%s\n", env.Str("EDITOR", "unknown"))
	fmt.Printf("$EMAIL:\t\t\t%s\n", env.Str("EMAIL", "unknown"))
	fmt.Printf("$PWD:\t\t\t%s\n", env.Str("PWD", "unknown"))
	fmt.Printf("$SHLVL:\t\t\t%d\n", env.Int("SHLVL", 0))
}
