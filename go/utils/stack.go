package utils

// Stack of strings
type StackStrings struct {
	elems []string
	nr    int
}

func (s *StackStrings) Push(elem string) {
	s.elems = append(s.elems, elem)
	s.nr++
}

// Pop - get element if available as signaled by ok
func (s *StackStrings) Pop() (elem string, ok bool) {
	if s.nr == 0 {
		return "", false
	}
	elem = s.elems[s.nr-1]
	s.nr--
	s.elems = s.elems[:s.nr]
	return elem, true
}

func (s *StackStrings) IsEmpty() bool {
	return s.nr == 0
}

func (s *StackStrings) Depth() int {
	return s.nr
}

// Stack of ints
type StackInts struct {
	elems []int
	nr    int
}

func (s *StackInts) Push(elem int) {
	s.elems = append(s.elems, elem)
	s.nr++
}

// Pop - get element if available as signaled by ok
func (s *StackInts) Pop() (elem int, ok bool) {
	if s.nr == 0 {
		return 0, false
	}
	elem = s.elems[s.nr-1]
	s.nr--
	s.elems = s.elems[:s.nr]
	return elem, true
}

func (s *StackInts) IsEmpty() bool {
	return s.nr == 0
}

func (s *StackInts) Depth() int {
	return s.nr
}
