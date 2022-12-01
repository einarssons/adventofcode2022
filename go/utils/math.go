package utils

func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinMaxInts(numbers []int) (int, int) {
	min, max := numbers[0], numbers[0]
	for _, nr := range numbers {
		if nr > max {
			max = nr
		}
		if nr < min {
			min = nr
		}
	}
	return min, max
}

func Min(numbers []int) int {
	minNr := 1 << 40
	for _, nr := range numbers {
		if nr < minNr {
			minNr = nr
		}
	}
	return minNr
}

func Max(numbers []int) int {
	maxNr := -(1 << 40)
	for _, nr := range numbers {
		if nr > maxNr {
			maxNr = nr
		}
	}
	return maxNr
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Triangle(nr int) int {
	return nr * (nr + 1) / 2
}

// GCDuint64 - greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}
