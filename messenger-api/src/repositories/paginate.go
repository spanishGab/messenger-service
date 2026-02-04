package repositories

import (
	"fmt"
)

type PaginatedResult struct {
	Content []any `json:"content"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	page uint8
	pageSize uint8
}

func (p Pagination) PaginateInMemory(content any) (*PaginatedResult, error) {
	parsedContent, isArray := content.([]any)

	if !isArray {
		return nil, fmt.Errorf("content is  not an array")
	}

	if p.page <= 0 {
		return nil, fmt.Errorf("paginateInMemory: 'page' must be greater than zero")
	}

	if p.pageSize <= 0 {
		return nil, fmt.Errorf("paginateInMemory: 'pageSize' must be greater than zero")
	}

	contentSize := uint8(len(parsedContent))

	start := (p.page - 1) * p.pageSize
	if start >= contentSize {
		return &PaginatedResult{
			Content: []any{},
			Pagination: Pagination{
				page: p.page,
				pageSize: p.pageSize,
			},
		}, nil
	}

	end := start + p.pageSize
	if end > contentSize {
		end = contentSize
	}

	var result PaginatedResult
	for i := start; i < end; i++ {
		result.Content = append(result.Content, parsedContent[i])
	}

	return &result, nil
}
