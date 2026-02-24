package handlers

type Handler func(Command) (string, error)

type OutputFormat string

const (
	JSONFormat  OutputFormat = "json"
	Unformatted OutputFormat = "unformatted"
)

func (of OutputFormat) IsValid() bool {
	switch of {
	case JSONFormat, Unformatted:
		return true
	default:
		return false
	}
}

func (of OutputFormat) String() string {
	return string(of)
}
