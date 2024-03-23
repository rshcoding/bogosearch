package bogosearch

import "math/rand/v2"

type indexedValue struct {
	value int
	index int
}

// Search attempts to find the target integer within the provided array.
// If the target is found, Search returns the original index of the target within arr and true;
// If the target is not found (i.e., arr becomes empty through reduction), it returns 0 and false.
//
// Parameters:
// - arr: The array of integers to be searched.
// - target: The integer value to search for within arr.
//
// Returns:
// - int: The original index of the target in the array if found; otherwise, returns 0.
// - bool: True if the target is found within the array; false if the target is not found.
func Search(arr []int, target int) (int, bool) {
	indexedArr := make([]indexedValue, len(arr))
	for i, val := range arr {
		indexedArr[i] = indexedValue{value: val, index: i}
	}
	return bogosearch(indexedArr, target)
}

func bogosearch(arr []indexedValue, target int) (int, bool) {
	if len(arr) == 0 {
		return 0, false
	}
	guess := rand.IntN(len(arr))
	if arr[guess].value == target {
		return arr[guess].index, true
	} else {
		newArr := make([]indexedValue, 0, len(arr)-1)
		newArr = append(newArr, arr[:guess]...)
		newArr = append(newArr, arr[guess+1:]...)
		return bogosearch(newArr, target)
	}
}
