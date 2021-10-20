package utils

// RemoveEmptyStrings remove empty strings from a slice
func RemoveEmptyStrings(s []string) []string {
	var slice []string
	for _, str := range s {
		if str != "" {
			slice = append(slice, str)
		}
	}
	return slice
}

// Utility function to check if `v` is present on array `s`
func ContainsString(s []string, v string) bool {
	for _, val := range s {
		if val == v {
			return true
		}
	}

	return false
}
