package utils

type Level int

const (
	Info Level = iota
	Success
	Warning
	Danger
)

type Messages struct {
	Message []string
	Type    Level
}

var (
	MessagesSuccess *Messages = &Messages{
		Message: make([]string, 0),
		Type:    Success,
	}
	MessagesDanger *Messages = &Messages{
		Message: make([]string, 0),
		Type:    Danger,
	}
	MessagesInfo *Messages = &Messages{
		Message: make([]string, 0),
		Type:    Danger,
	}
)
