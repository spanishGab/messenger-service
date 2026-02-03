package repositories

import (
	"fmt"
	"messenger-api/src/entities"
)

type PaginatedResults struct {
	Content any `json:"content"`
	Pagination any `json:"pagination"`
}

type Pagination struct {
	page uint8
	pageSize uint8
}

func (p Pagination) paginateInMemory(content []entities.Message) ([]entities.Message, error) {
	if p.page <= 0 {
		return nil, fmt.Errorf("paginateInMemory: 'page' must be greater than zero")
	}

	if p.pageSize <= 0 {
		return nil, fmt.Errorf("paginateInMemory: 'pageSize' must be greater than zero")
	}

	contentSize := uint8(len(content))

	start := (p.page - 1) * p.pageSize
	if start >= contentSize {
		return []entities.Message{}, nil
	}

	end := start + p.pageSize
	if end > contentSize {
		end = contentSize
	}

	var result []entities.Message
	for i := start; i < end; i++ {
		result = append(result, content[i])
	}

	return result, nil
}
