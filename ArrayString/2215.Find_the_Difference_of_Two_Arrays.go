package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {
	//create map of both the arrays
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	// mark elements present in each as true
	for _, num := range nums1 {
		map1[num] = true
	}

	for _, num := range nums2 {
		map2[num] = true
	}

	//Create two empty arrays to store unique elements of both maps
	uniqueN1 := []int{}
	uniqueN2 := []int{}

	//Match the Key of map1 in map2 to check the presence
	//We marked the elemts present as true so if key is not true(!true)
	//store in uniqueN1 and N2 respectively

	for key := range map1 {
		if !map2[key] {
			uniqueN1 = append(uniqueN1, key)
		}
	}

	for key := range map2 {
		if !map1[key] {
			uniqueN2 = append(uniqueN2, key)
		}
	}
	return [][]int{uniqueN1, uniqueN2}
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	result := findDifference(nums1, nums2)
	fmt.Print(result)

}
