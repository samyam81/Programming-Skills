package main

func judgeCircle(moves string) bool {
    right:=0
	left:=0

	for _, kobbie := range moves {
		if kobbie=='U' {
			right++
		} else if kobbie=='D' {
			right--
		} else if kobbie=='R' {
			left++
		} else {
			left--
		}
	}

	if left==0 && right==0 {
		return true
	}

	return false
}

// Runtime:= 4ms Beats 72.00%
// Memory:= 3.44MB Beats 26.67%