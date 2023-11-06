package day_25

import "strconv"

func snafuToDecimal(s string) int {
	var positive, negative int

	// Reverse the string
	digits := reverseString(s)

	// Build the positive and negative numbers
	for i, r := range digits {
		switch r {
		case '2':
			positive += 2 * power(5, i)
		case '1':
			positive += power(5, i)
		case '=':
			negative += 2 * power(5, i)
		case '-':
			negative += power(5, i)
		default:
			continue
		}
	}

	// Subtract the negative number from the positive number
	return positive - negative
}

func reverseString(s string) string {
	var out string

	for _, r := range s {
		out = string(r) + out
	}

	return out
}

// Quick and dirty power function for ints - far from optimal
func power(base, power int) int {
	out := 1

	for i := power; i > 0; i-- {
		out *= base
	}

	return out
}

func decimalToSnafu(i int) string {
	var out string

	// Convert the decimal number to base-5
	base5 := decimalToBase5(i)

	// Reverse it
	digits := reverseString(base5)

	// Build the snafu number right-to-left
	var carry bool
	for _, r := range digits {
		if carry {
			switch r {
			case '4':
				out = "0" + out
			case '3':
				out = "-" + out
			case '2':
				out = "=" + out
			case '1':
				out = "2" + out
				carry = false
			case '0':
				out = "1" + out
				carry = false
			}
		} else {
			switch r {
			case '4':
				out = "-" + out
				carry = true
			case '3':
				out = "=" + out
				carry = true
			default:
				out = string(r) + out
			}
		}
	}

	if carry {
		out = "1" + out
	}

	return out
}

func decimalToBase5(i int) string {
	// Figure out the biggest power we'll need
	var biggest int

	for i > power(5, biggest)-1 {
		biggest += 1
	}

	biggest -= 1

	// Build the number left-to-right
	num := i
	var out string
	for pow := biggest; pow >= 0; pow-- {
		out = out + strconv.Itoa(num/power(5, pow))
		num = num % power(5, pow)
	}

	return out
}
