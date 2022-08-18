package sum

// Sum of slice
func Sum(slice []int) (total int) {
	for _, element := range slice {
		total += element
	}
	return
}

// All many slices, returning slice of results
func All(slices ...[]int) []int {
	totals := make([]int, len(slices))
	for i, slice := range slices {
		totals[i] = Sum(slice)
	}
	return totals
}
