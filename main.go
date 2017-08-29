package main

import (
	"fmt"

	"github.com/waggertron/go_data_structures/set"
)

func main() {
	s1 := set.IntSet(1, 2, 3)
	s2 := set.IntSet(3, 4, 5)
	s3 := s1.Union(&s2)
	fmt.Println(s3.SortedItems())
}
