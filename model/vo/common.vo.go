package vo

type PaginationResult struct {
	Total       int         `json:"total"`
	Pages       int         `json:"pages"`
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	Records     interface{} `json:"records"`
}
