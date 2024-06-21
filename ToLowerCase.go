package main


func toLowerCase(s string) string {

	chars:=[]rune(s)
	for i, char := range chars{
		if char>='A' && char<='Z' {
			chars[i]=char+32
		}		
	}

	return string(chars)
}

// Runtime:= 1ms Beats 72.25%
// Memory:= 2.36MB Beats 30.06%