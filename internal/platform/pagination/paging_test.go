package pagination_test

import (
	"net/url"
	"testing"

	"github.com/jbelmont/api-workshop/internal/platform/pagination"
)

func TestPopulatePaging(t *testing.T) {
	url := url.URL{
		Scheme:   "http",
		Host:     "localhost:8080",
		Path:     "/api/v1/heroes",
		RawQuery: "pageSize=50&pageIndex=1",
	}

	paging := pagination.PopulatePaging(url, 50, 100)
	if paging.Index != 1 {
		t.Errorf("Expected %d, Actual %d", 1, paging.Index)
	}
	if paging.Size != 50 {
		t.Errorf("Expected %d, Actual %d", 50, paging.Size)
	}
}
