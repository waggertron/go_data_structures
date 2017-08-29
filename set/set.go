package set

import "sort"

type intSet struct {
	items map[int]bool
}

func IntSet(nums ...int) intSet {
	s := intSet{items: make(map[int]bool)}
	s.Add(nums...)
	return s
}

func (s *intSet) Add(nums ...int) {
	for _, num := range nums {
		s.items[num] = true
	}
}

func (s *intSet) Has(n int) bool {
	_, present := s.items[n]
	return present
}

func (s *intSet) Remove(n int) {
	delete(s.items, n)
}
func (s *intSet) Intersection(s2 *intSet) intSet {
	inter := IntSet()
	var smaller, bigger *intSet
	if len(s.items) <= len(s2.items) {
		smaller, bigger = s, s2
	} else {
		smaller, bigger = s2, s
	}
	for num := range smaller.items {
		if bigger.Has(num) {
			inter.Add(num)
		}
	}
	return inter
}

func (s *intSet) Union(s2 *intSet) intSet {
	union := IntSet()
	for num := range s.items {
		union.Add(num)
	}
	for num := range s2.items {
		union.Add(num)
	}
	return union
}

func (s *intSet) Items() []int {
	items := make([]int, 0, len(s.items))
	for num := range s.items {
		items = append(items, num)
	}
	return items
}
func (s *intSet) SortedItems() []int {
	items := make([]int, 0, len(s.items))
	for num := range s.items {
		items = append(items, num)
	}
	sort.Ints(items)
	return items
}

// func (s *intSet) add(nums ...int) {
// 	for _, num := range nums {
// 		s.items[num] = true
// 	}
// }
