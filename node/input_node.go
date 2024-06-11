package node

import "github.com/ninomae42/node_interface/domain"

type (
	inputNode struct {
		NodeID domain.ID

		Label domain.String

		Required bool
	}

	TextInputNode struct {
		inputNode

		MaxLength int
	}

	NumberInputNode struct {
		inputNode

		Unit *domain.String
	}
)

func (i inputNode) ID() domain.ID { return i.NodeID }
func (inputNode) Type() Type      { return NodeTypeInput }

func NewTextInputNode(label domain.String, required bool, maxLength int) *TextInputNode {
	return &TextInputNode{
		inputNode: inputNode{
			NodeID:   domain.NewID(),
			Label:    label,
			Required: required,
		},
		MaxLength: maxLength,
	}
}
