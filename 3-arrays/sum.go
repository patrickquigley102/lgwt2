package sum

// Sum of slice
func Sum(slice []int) (total int) {
	for _, element := range slice {
		total += element
	}
	return
}
