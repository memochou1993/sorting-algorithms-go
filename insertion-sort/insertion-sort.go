package main

// InsertionSort func
func InsertionSort(items []int) []int {
	for i := 1; i < len(items); i++ {
		temp := items[i]

		j := i - 1
		for ; j >= 0 && temp < items[j]; j-- {
			items[j+1] = items[j]
		}

		items[j+1] = temp
	}

	return items
}
