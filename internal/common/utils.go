package common

import "strconv"

func ParseUint(s string) (uint, error) {
	id64, err := strconv.ParseUint(s, 10, 64)
	return uint(id64), err
}
