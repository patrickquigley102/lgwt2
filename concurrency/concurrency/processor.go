package concurrency

import "fmt"

const (
	okString    = "Status OK"
	notOkString = "Status Not OK"
	hundred     = 100.0
)

// Process expresses boolean frequency as percentage.
func Process(dataSet map[string]bool) map[string]string {
	results := map[string]string{}
	total := float64(len(dataSet))
	var success float64

	for _, b := range dataSet {
		if b {
			success += 1.0
		}
	}

	success = (success / total) * hundred
	results[okString] = fmtPercentage(success)
	results[notOkString] = fmtPercentage(hundred - success)

	return results
}

func fmtPercentage(n float64) string {
	return fmt.Sprintf("%.0f", n) + "%"
}
