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
