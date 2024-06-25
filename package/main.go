// every package should start with package name
// there can be multiple packages main
// but they're considered as one
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(100)

	fmt.Println(n)
}
