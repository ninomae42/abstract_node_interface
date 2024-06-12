package node

import "github.com/ninomae42/node_interface/domain"

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
