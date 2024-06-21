package main

func mergeAlternately(word1 string, word2 string) string {
    var merged string=""
	maxLength := len(word1)
	if len(word2) > maxLength {
		maxLength = len(word2)
	}

	for i := 0; i < maxLength; i++ {
		if i< len(word1) {
			merged+=string(word1[i])
		}
		if i< len(word2) {
			merged+=string(word2[i])
		}

	}
	return merged
}

// Runtime:= 0ms Beats 100%
// Memory:= 3.06MB Beats 21.10%