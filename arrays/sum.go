// Package sum for learnings
package sum

// Sum of slice.
func Sum(slice []int) int {
	total := 0
	for _, element := range slice {
		total += element
	}

	return total
}

// All many slices, returning slice of results.
func All(slices ...[]int) []int {
	totals := make([]int, len(slices))
	for i, slice := range slices {
		totals[i] = Sum(slice)
	}

	return totals
}
