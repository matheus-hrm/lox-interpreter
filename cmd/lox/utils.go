package main

import (
	"fmt"
	"strconv"
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
		formattedValue := strconv.FormatFloat(number, 'f', -1, 64)
		return strings.TrimRight(formattedValue, "0")
	}
}
