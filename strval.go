// Package strval provides functionality for validation of UTF-8 encoded strings.
package strval

import (
	"strings"
)

// Validate validates if assertions asserts, for string s are true.
func Validate(s string, asserts ...AssertFunc) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}

	for _, a := range asserts {
		if ok := a(s); !ok {
			return false
		}
	}
	return true
}

// ValidateSubstr validates if assertions asserts for substring s[low:high] are true.
func ValidateSubstr(low, high uint, s string, asserts ...AssertFunc) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}

	for _, a := range asserts {
		var (
			ss      string
			sLow, i uint
		)
		for sHigh := range s {
			if i == low {
				sLow = uint(sHigh)
			}
			if i == high {
				ss = s[sLow:sHigh]
				goto assert
			}
			i++
		}
		ss = s[sLow:]

	assert:
		if ok := a(ss); !ok {
			return false
		}
	}
	return true
}
