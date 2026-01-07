package handlers

type CommandType string

const (
	Create CommandType = "create"
	Update CommandType = "update"
	Delete CommandType = "delete"
	List   CommandType = "list"
)

type Command struct {
	Type CommandType
	Data map[string]string
}
