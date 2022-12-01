package utils

// Set - mathematical set with operations
// Empty struct requires zero bytes so is more efficient than bool
type Set map[string]struct{}

// CreateSet - create an empty set
func CreateSet() Set {
	return Set(make(map[string]struct{}))
}

// Contains - check if elem in set
func (s Set) Contains(elem string) bool {
	_, ok := s[elem]
	return ok
}

// Add - add elem to set
func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

// Remove - remove elem from set (does not need to be in set)
func (s Set) Remove(elem string) {
	delete(s, elem)
}

// Extend - extend set s with all elements in other (the result is union)
func (s Set) Extend(other Set) {
	for k := range other {
		s[k] = struct{}{}
	}
}

// Subtract - remove all elements from s that are in other
func (s Set) Subtract(other Set) {
	for k := range other {
		_, ok := s[k]
		if ok {
			delete(s, k)
		}
	}
}

// Intersect - only keep elements in s which are also in other
func (s Set) Intersect(other Set) {
	deleteList := make([]string, 0, len(s))
	for k := range s {
		_, inOther := other[k]
		if !inOther {
			deleteList = append(deleteList, k)
		}
	}
	for _, k := range deleteList {
		delete(s, k)
	}
}

// SetInts - mathematical set with operations
// Empty struct requires zero bytes so is more efficient than bool
type SetInts map[int]struct{}

// CreateSetInts - create an empty set
func CreateSetInts() SetInts {
	return SetInts(make(map[int]struct{}))
}

// Contains - check if elem in set
func (s SetInts) Contains(elem int) bool {
	_, ok := s[elem]
	return ok
}

// Add - add elem to set
func (s SetInts) Add(elem int) {
	s[elem] = struct{}{}
}

// Remove - remove elem from set (does not need to be in set)
func (s SetInts) Remove(elem int) {
	delete(s, elem)
}

// Extend - extend set s with all elements in other (the result is union)
func (s SetInts) Extend(other SetInts) {
	for k := range other {
		s[k] = struct{}{}
	}
}

// Subtract - remove all elements from s that are in other
func (s SetInts) Subtract(other SetInts) {
	for k := range other {
		_, ok := s[k]
		if ok {
			delete(s, k)
		}
	}
}

// Intersect - only keep elements in s which are also in other
func (s SetInts) Intersect(other SetInts) {
	deleteList := make([]int, 0, len(s))
	for k := range s {
		_, inOther := other[k]
		if !inOther {
			deleteList = append(deleteList, k)
		}
	}
	for _, k := range deleteList {
		delete(s, k)
	}
}
