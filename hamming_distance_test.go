package hamming_distance

import "testing"

func TestHammingDistance(t *testing.T) {
	a := 1
	b := 3
	distance := Int(a, b)
	t.Log(distance)
	// Output:
	// 1
}

func TestString(t *testing.T) {
	s1 := "abcd"
	s2 := "abcf"
	distance := String(s1, s2)
	t.Log(distance)
	// Output:
	// 1
}
