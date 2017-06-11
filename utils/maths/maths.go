//
// Package maths contains helpers to the mathematical calculations
//
package maths

// Round returns the rounded value of val to precision 0
//
// Round(3.4);       // 3
// Round(3.5);       // 4
// Round(3.6);       // 4
// Round(1.95583);   // 2
// Round(1241757.0); // 1241757
// Round(-4.27)      // -4
// Round(-0.67)      // -1
//
func Round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}

	return int(val + 0.5)
}
