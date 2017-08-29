package main

import (
	"fmt"

	"github.com/waggertron/set"
)

func main() {
	s1 := set.IntSet(1, 2, 3)
	s2 := set.IntSet(3, 4, 5)
	fmt.Println(s1.Intersection(&s2))
}
