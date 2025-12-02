package common

import "strconv"

func GetPageSize(pageStr, sizeStr string) (page, size int) {
	page, _ = strconv.Atoi(pageStr)
	size, _ = strconv.Atoi(sizeStr)

	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 100 {
		size = 10
	}
	return
}
