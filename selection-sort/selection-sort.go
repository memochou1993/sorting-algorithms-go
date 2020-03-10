package main

// SelectionSort func
func SelectionSort(items []int) []int {
	for i := 0; i < len(items); i++ {
		min := i

		for j := i + 1; j < len(items); j++ {
			if items[min] > items[j] {
				min = j
			}
		}

		temp := items[min]
		items[min] = items[i]
		items[i] = temp
	}

	return items
}
