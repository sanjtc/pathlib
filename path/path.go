package path

import "strings"

// ConvertBackslashToSlash
// Convert "\\" to "/".
func ConvertBackslashToSlash(s string) string {
	ss := strings.Split(s, "\\")
	res := ""

	for i, t := range ss {
		res += t

		if i < len(ss)-1 {
			res += "/"
		}
	}

	return res
}
