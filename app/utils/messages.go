package utils

type Level string

const (
	Info    Level = "info"
	Success Level = "success"
	Warning Level = "warning"
	Danger  Level = "danger"
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

func (self *Messages) AddMessage(msg string) {
	switch self.Type {
	case "info":
		MessagesInfo.Message = append(MessagesInfo.Message, msg)
	case "danger":
		MessagesDanger.Message = append(MessagesDanger.Message, msg)
	case "success":
		MessagesSuccess.Message = append(MessagesSuccess.Message, msg)
	}
}

func (self *Messages) ClearMessage() {
	self.Message = make([]string, 0)
}
