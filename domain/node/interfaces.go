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
