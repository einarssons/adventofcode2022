package utils

func ReverseStrings(s []string) {
	l := len(s)
	nrSwaps := l / 2
	for i := 0; i < nrSwaps; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}

func ReverseInts(s []int) {
	l := len(s)
	nrSwaps := l / 2
	for i := 0; i < nrSwaps; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}
