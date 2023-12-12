package pagination

import "math"

func MyPagination(limit int, totalRow int, currentPage int) (pagination map[string]interface{}) {
	var totalPage float64 = 0

	if limit > 0 {
		totalPage = math.Ceil(float64(totalRow) / float64(limit))
	}

	pagination = map[string]interface{}{
		"perPages":     limit,
		"numOfPages":   totalPage,
		"numOfResults": totalRow,
		"currentPage":  currentPage,
	}

	return
}
