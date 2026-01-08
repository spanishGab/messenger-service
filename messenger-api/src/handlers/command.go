package handlers

type CommandType string 

const (
	Create CommandType = "CREATE"
	List CommandType = "LIST"
	ListByFilters CommandType = "LIST_BY_FILTERS"
	Update CommandType = "UPDATE"
	Delete CommandType = "DELETE"
)

type DateRangeDTO struct {
	Start string
	End  string
}

type Command struct {
	Type CommandType
	Data map[string]string
}
