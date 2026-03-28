package repositories

import (
	"fmt"
)

type PaginatedResult[T any] struct {
	Items []T `json:"items"`
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

func PaginateInMemory[T any](items []T, pagination Pagination) (*PaginatedResult[T], error) {
	contentSize := uint8(len(items))

	start := (pagination.Page - 1) * pagination.PageSize
	if start >= contentSize {
		return &PaginatedResult[T]{
			Items: []T{},
			Pagination: pagination,
		}, nil
	}

	end := start + pagination.PageSize
	if end > contentSize {
		end = contentSize
	}

	result := PaginatedResult[T]{Pagination: pagination}
	for i := start; i < end; i++ {
		result.Items = append(result.Items, items[i])
	}

	return &result, nil
}
