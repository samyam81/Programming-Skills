package main

import "strconv"

func calPoints(operations []string) int {
	size := 0
	scores := make([]int, len(operations))

	for _, op := range operations {
		if op == "+" {
			if size >= 2 {
				scores[size] = scores[size-1] + scores[size-2]
				size++
			}
		} else if op == "D" {
			if size >= 1 {
				scores[size] = 2 * scores[size-1]
				size++
			}
		} else if op == "C" {
			if size > 0 {
				size--
			}
		} else {
			num, err := strconv.Atoi(op)
			if err == nil {
				scores[size] = num
				size++
			}
		}
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum += scores[i]
	}

	return sum
}

// Rumtime:= 2ms Beats 76.14%
// Memory:= 2.72MB Beats 68.53%