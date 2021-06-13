package conf

import (
	"fmt"
	"os"
	"strconv"
)

const confPrefix = "RAINY_"

// GetString returns the value of a configuration setting as a string
func GetString(val string) string {
	return os.Getenv(fmt.Sprintf("%s%s", confPrefix, val))
}

// GetInt returns the value of a configuration setting as an int32, or 0 if it is invalid
func GetInt(val string) int {
	ix, err := strconv.Atoi(os.Getenv(fmt.Sprintf("%s%s", confPrefix, val)))

	if err != nil {
		return 0
	}

	return ix
}

// GetUInt64 returns the value of a configuration setting as an uint64, or 0 if it is invalid
func GetUInt64(val string) uint64 {
	ix, err := strconv.ParseUint(os.Getenv(fmt.Sprintf("%s%s", confPrefix, val)), 10, 64)

	if err != nil {
		return 0
	}

	return ix
}
