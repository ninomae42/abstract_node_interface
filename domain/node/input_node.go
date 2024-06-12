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

	TextInputNode struct {
		NodeID domain.ID

		Label domain.String

		Required bool

		MaxLength int
	}

	NumberInputNode struct {
		NodeID domain.ID

		Label domain.String

		Required bool

		Unit *domain.String
	}
)

var _ InputNode = (*TextInputNode)(nil)

func (n *TextInputNode) ID() domain.ID        { return n.NodeID }
func (n *TextInputNode) Type() Type           { return NodeTypeInput }
func (n *TextInputNode) InputType() InputType { return InputTypeText }

var _ InputNode = (*NumberInputNode)(nil)

func (n *NumberInputNode) ID() domain.ID        { return n.NodeID }
func (n *NumberInputNode) Type() Type           { return NodeTypeInput }
func (n *NumberInputNode) InputType() InputType { return InputTypeNumber }

func NewTextInputNode(label domain.String, required bool, maxLength int) *TextInputNode {
	return &TextInputNode{
		NodeID:    domain.NewID(),
		Label:     label,
		Required:  required,
		MaxLength: maxLength,
	}
}
