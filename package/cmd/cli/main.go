// every package should start with package name
// there can be multiple packages main
// but they're considered as one
package main

import (
	"fmt"
	"package/internal/random"
)

func main() {
	n := random.RandomNumber()

	fmt.Println(n)
}
