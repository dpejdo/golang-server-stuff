// every package should start with package name
// there can be multiple packages main
// but they're considered as one
package main

import (
	"package/internal/random"

	"github.com/fatih/color"
)

func main() {

	blue := color.New(color.FgBlue)

	blue.Printf("lucky number: %d", random.RandomNumber())
}
