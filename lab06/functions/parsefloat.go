package functions

import (
	"strconv"
	"strings"
)

func ParseFloat32(s string) (float32, error) {
	cleanedString := strings.ReplaceAll(s, "$", "")
	val, err := strconv.ParseFloat(cleanedString, 32)
	if err != nil {
		return 0, err
	}
	return float32(val), nil
}