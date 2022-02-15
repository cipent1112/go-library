package conv

import "strings"

func Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}
