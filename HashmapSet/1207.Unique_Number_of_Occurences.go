
func uniqueOccurences(arr []int) bool {
	occurences := make(map[int]int)
	unique := make(map[int]bool)

	for _, val := range arr {
		occurences[val]++
	}

	for _, count := range occurences {
		if unique[count] {
			return false
		}
		unique[count] = true
	}
	return true
}
