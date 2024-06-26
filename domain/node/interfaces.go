package node

import "github.com/ninomae42/node_interface/domain"

type (
	// Node All node types implement the Node interface
	Node interface {
		ID() domain.ID // ID returns the node's ID
		Type() Type    // Type returns the node's type
	}

	// InputNode All input node types implement the InputNode interface
	InputNode interface {
		Node
		InputType() InputType // InputType returns the input node's input type
	}
)

type (
	// Type is a type of node
	Type string
	// InputType is a type of input node
	InputType string
)

const (
	// NodeTypeGroup is a group node type
	NodeTypeGroup Type = "NODE_GROUP"

	// NodeTypeInput is an input node type
	NodeTypeInput Type = "NODE_INPUT"

	// NodeTypePreview is a preview node type
	NodeTypePreview Type = "NODE_PREVIEW"
)

const (
	InputTypeText   InputType = "INPUT_TEXT"
	InputTypeNumber InputType = "INPUT_NUMBER"
)
