// Package numerals is for numerals.
package numerals

// ConvertToNumerals converts an integer to sting of Roman numerals.
func ConvertToNumerals(num int) (string, error) {
	var result string
	for num > 0 {
		result += "I"
		num--
	}

	return result, nil
}
