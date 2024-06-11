package fixture

import (
	"github.com/ninomae42/node_interface/domain"
	"github.com/ninomae42/node_interface/node"
)

func GroupNode(opts ...func(*node.GroupNode)) *node.GroupNode {
	n := &node.GroupNode{
		NodeID: domain.NewID(),
		Label:  domain.NewString(),
		Level:  node.GroupLevel1,
		Nodes:  make([]*node.Node, 0),
	}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

func GroupNodeWithID(id domain.ID) func(*node.GroupNode) {
	return func(n *node.GroupNode) {
		n.NodeID = id
	}
}

func GroupNodeWithLabel(label domain.String) func(*node.GroupNode) {
	return func(n *node.GroupNode) {
		n.Label = label
	}
}

func GroupNodeWithLevel(level node.GroupLevel) func(*node.GroupNode) {
	return func(n *node.GroupNode) {
		n.Level = level
	}
}

func GroupNodeWithNodes(nodes ...node.Node) func(*node.GroupNode) {
	return func(n *node.GroupNode) {
		for _, node := range nodes {
			n.Nodes = append(n.Nodes, &node)
		}
	}
}

func TextInputNode(opts ...func(*node.TextInputNode)) *node.TextInputNode {
	n := node.NewTextInputNode(domain.NewString(), false, 0)
	for _, opt := range opts {
		opt(n)
	}
	return n
}

func TextInputNodeWithID(id domain.ID) func(*node.TextInputNode) {
	return func(n *node.TextInputNode) {
		n.NodeID = id
	}
}

func TextInputNodeWithLabel(label domain.String) func(*node.TextInputNode) {
	return func(n *node.TextInputNode) {
		n.Label = label
	}
}

func TextInputNodeWithMaxLength(maxLength int) func(*node.TextInputNode) {
	return func(n *node.TextInputNode) {
		n.MaxLength = maxLength
	}
}
