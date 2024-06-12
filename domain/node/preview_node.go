package node

import "github.com/ninomae42/node_interface/domain"

type PreviewNode struct {
	NodeID domain.ID

	Text domain.String
}

var _ Node = (*PreviewNode)(nil)

func (n *PreviewNode) ID() domain.ID { return n.NodeID }
func (*PreviewNode) Type() Type      { return NodeTypePreview }
