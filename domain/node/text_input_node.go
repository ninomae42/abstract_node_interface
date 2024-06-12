package node

import "github.com/ninomae42/node_interface/domain"

type (
	TextInputNode struct {
		NodeID domain.ID

		Label domain.String

		Required bool

		MaxLength int
	}
)

var _ InputNode = (*TextInputNode)(nil)

func (n *TextInputNode) ID() domain.ID        { return n.NodeID }
func (n *TextInputNode) Type() Type           { return NodeTypeInput }
func (n *TextInputNode) InputType() InputType { return InputTypeText }

func NewTextInputNode(label domain.String, required bool, maxLength int) *TextInputNode {
	return &TextInputNode{
		NodeID:    domain.NewID(),
		Label:     label,
		Required:  required,
		MaxLength: maxLength,
	}
}
