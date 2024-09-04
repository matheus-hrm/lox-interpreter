package main

import (
	"fmt"
	"strings"
)

func formatFloat(value string, number float64) string {
	hasFractionalPart := strings.Contains(value, ".")

	zeroFraction := func(value string) bool {
		if i := strings.Index(value, "."); i != -1 {
			for _, digit := range value[i+1:] {
				if digit != '0' {
					return false
				}
			}
		}
		return true
	}

	if !hasFractionalPart || zeroFraction(value) {
		return fmt.Sprintf("%.1f", number)
	} else {
		decimalCount := len(value) - strings.Index(value, ".") - 1
		format := fmt.Sprintf("%%.%df", decimalCount)
		return fmt.Sprintf(format, number)
	}
}
