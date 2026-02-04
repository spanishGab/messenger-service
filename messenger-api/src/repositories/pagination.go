package repositories

import (
	"fmt"
)

type PaginatedResult struct {
	Content []any `json:"content"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page uint8
	PageSize uint8
}

func NewPagination(page uint8, pageSize uint8) (*Pagination, error) {
	if page <= 0 {
		return nil, fmt.Errorf("paginateInMemory: 'page' must be greater than zero")
	}

	if pageSize <= 0 {
		return nil, fmt.Errorf("paginateInMemory: 'pageSize' must be greater than zero")
	}

	return &Pagination{
		Page: page,
		PageSize: pageSize,
	}, nil
}

func (p Pagination) PaginateInMemory(content any) (*PaginatedResult, error) {
	parsedContent, isArray := content.([]any)

	if !isArray {
		return nil, fmt.Errorf("content is  not an array")
	}

	contentSize := uint8(len(parsedContent))

	start := (p.Page - 1) * p.PageSize
	if start >= contentSize {
		return &PaginatedResult{
			Content: []any{},
			Pagination: Pagination{
				Page: p.Page,
				PageSize: p.PageSize,
			},
		}, nil
	}

	end := start + p.PageSize
	if end > contentSize {
		end = contentSize
	}

	var result PaginatedResult
	for i := start; i < end; i++ {
		result.Content = append(result.Content, parsedContent[i])
	}

	return &result, nil
}
