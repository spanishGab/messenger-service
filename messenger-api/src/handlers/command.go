package handlers

type CommandType string 

const (
	Create CommandType = "CREATE"
	ListById CommandType = "LIST_BY_ID"
	List CommandType = "LIST"
	Update CommandType = "UPDATE"
	Delete CommandType = "DELETE"
)

type Command struct {
	Type CommandType
	Data map[string]string
}
