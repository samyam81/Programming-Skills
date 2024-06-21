package main

func maximumWealth(accounts [][]int) int {
    maxWealth:=0
	for i := 0; i < len(accounts); i++ {
		wealth:=0
		for j := 0; j < len(accounts[i]); j++ {
			 wealth += accounts[i][j] 
		}

	maxWealth=int(max(float64(wealth),float64(maxWealth)))
	}

	return maxWealth
}

// Runtime:= 4ms Beats 31.78%
// Memory:= 2.93MB Beats 31.09%