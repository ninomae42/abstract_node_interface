package node

type InputType string

const (
	InputTypeText   InputType = "INPUT_TEXT"
	InputTypeNumber InputType = "INPUT_NUMBER"
)

type (
	InputNode interface {
		Node
		InputType() InputType
	}
)
