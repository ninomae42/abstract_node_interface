package node

import "github.com/ninomae42/node_interface/domain"

type (
	Node interface {
		ID() domain.ID
		Type() Type
	}

	InputNode interface {
		Node
		InputType() InputType
	}
)

type (
	Type      string
	InputType string
)

const (
	NodeTypeGroup Type = "NODE_GROUP"
	NodeTypeInput Type = "NODE_INPUT"
)

const (
	InputTypeText   InputType = "INPUT_TEXT"
	InputTypeNumber InputType = "INPUT_NUMBER"
)
