package main

func diagonalSum(mat [][]int) int {
    i,j,s:=0,0,0
	
	for i<len(mat){
		s+=mat[i][j]
		i++
		j++
	}
	
	j=len(mat)-1
	i=0

	for i<len(mat){
		if i!=j {
			s+=mat[i][j]	
		}
		i++;
        j--; 
	}

	return s
}

// Runtime:=10ms Beats 54.36%
// Memory:= 4.62MB Beats 14.77%