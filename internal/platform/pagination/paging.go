package pagination

import (
	"net/url"
	"strconv"
)

// Paging holds information about paging such as as the pageIndex etc.
type Paging struct {
	Index  int
	Size   int
	Sort   []string
	Select []string
}

// PopulatePaging extracts paging data from the URL into Paging.
// A max size should be set for performance concerns.
func PopulatePaging(url url.URL, defaultSize int, maxSize int) Paging {
	q := url.Query()

	paging := Paging{
		Size:   defaultSize,
		Sort:   q["sort"],
		Select: q["select"],
	}
	if size, err := strconv.Atoi(q.Get("pageSize")); err == nil && size > 0 {
		paging.Size = size
		if paging.Size > maxSize {
			paging.Size = maxSize
		}
	}

	if index, err := strconv.Atoi(q.Get("pageIndex")); err == nil && index > 0 {
		paging.Index = index
	}

	return paging
}
