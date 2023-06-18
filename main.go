package main

import (
	"fmt"
	"github.com/aryaOlya/algorithm/binarysearch"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	fmt.Println(binarysearch.BinarySearch(a, 16))

}
