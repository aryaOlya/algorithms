package binarysearch

import "fmt"

func BinarySearch(list []int, x int) string {
	r := -1 // not found
	start := 0
	end := len(list) - 1

	count := 0
	for start <= end {
		count++
		mid := (start + end) / 2
		if list[mid] == x {
			r = mid
			break
		} else if list[mid] < x {
			start = mid + 1
		} else if list[mid] > x {
			end = mid - 1
		}
	}

	result := fmt.Sprintf("item %v found in %v time try instead of %v",a[r],count,x)
	return result
}
