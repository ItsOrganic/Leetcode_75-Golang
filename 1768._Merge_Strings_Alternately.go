package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var merge string
	x, y := len(word1), len(word2)

	for i := 0; i < x || i < y; i++ {
		if x > i {
			merge += string(word1[i])
		}
		if y > i {
			merge += string(word2[i])
		}
	}
	return merge
}

func main() {
	word1, word2 := "abc", "pqr"
	result := mergeAlternately(word1, word2)
	fmt.Print(result)

}
