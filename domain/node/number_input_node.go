package node

import "github.com/ninomae42/node_interface/domain"

type (
	NumberInputNode struct {
		NodeID domain.ID

		Label domain.String

		Required bool

		Unit *domain.String
	}
)

var _ InputNode = (*NumberInputNode)(nil)

func (n *NumberInputNode) ID() domain.ID        { return n.NodeID }
func (n *NumberInputNode) Type() Type           { return NodeTypeInput }
func (n *NumberInputNode) InputType() InputType { return InputTypeNumber }
