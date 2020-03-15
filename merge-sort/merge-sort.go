package main

// MergeSort func
func MergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	half := len(items) / 2
	left := MergeSort(items[:half])
	right := MergeSort(items[half:])

	result := []int{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			item := left[0]
			left = left[1:]
			result = append(result, item)
		} else {
			item := right[0]
			right = right[1:]
			result = append(result, item)
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}
