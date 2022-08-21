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

// AllTails sums the slice of slices excluding the first (head) element.
func AllTails(slices ...[]int) []int {
	totals := make([]int, len(slices))
	for i, slice := range slices {
		if len(slice) == 0 {
			totals[i] = 0
		} else {
			totals[i] = Sum(slice[1:])
		}
	}

	return totals
}
