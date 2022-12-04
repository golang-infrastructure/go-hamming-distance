package hamming_distance

import (
	"fmt"
	"github.com/golang-infrastructure/go-gtypes"
)

func Int[T gtypes.Integer](x, y T) T {
	xor := x ^ y
	var distance T
	for xor != 0 {
		distance = distance + 1
		xor = xor & (xor - 1)
	}
	return distance
}

func String(s1, s2 string) int {
	distance, _ := StringE(s1, s2)
	return distance
}

func StringE(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, fmt.Errorf("can not compute hamming distance on diff length string %s and %s", s1, s2)
	}
	distance := 0
	for index := range s1 {
		if s1[index] != s2[index] {
			distance++
		}
	}
	return distance, nil
}
