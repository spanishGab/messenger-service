package handlers

type CommandType string 

const (
	Create CommandType = "create"
	List CommandType = "list"
	Update CommandType = "update"
	Delete CommandType = "delete"
)

type Command struct {
	Type CommandType
	Data map[string]string
}
