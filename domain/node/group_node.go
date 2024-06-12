package node

import (
	"github.com/ninomae42/node_interface/domain"
)

type GroupLevel string

const (
	GroupLevel1 GroupLevel = "GROUP_LEVEL_1"
	GroupLevel2 GroupLevel = "GROUP_LEVEL_2"
	GroupLevel3 GroupLevel = "GROUP_LEVEL_3"
	GroupLevel4 GroupLevel = "GROUP_LEVEL_4"
	GroupLevel5 GroupLevel = "GROUP_LEVEL_5"
)

type GroupNode struct {
	NodeID domain.ID
	Label  domain.String

	Level GroupLevel

	DirectionIsVertical bool

	Nodes []*Node
}

var _ Node = (*GroupNode)(nil)

func (n *GroupNode) ID() domain.ID { return n.NodeID }
func (*GroupNode) Type() Type      { return NodeTypeGroup }

func (g *GroupNode) AddNode(n []Node) {
	for _, node := range n {
		g.Nodes = append(g.Nodes, &node)
	}
}

func (g GroupNode) Flatten() ([]Node, error) {
	var nodes []Node
	nodes = append(nodes, &g)
	for _, n := range g.Nodes {
		switch n := (*n).(type) {
		case *GroupNode:
			subnodes, err := (*n).Flatten()
			if err != nil {
				return []Node{}, err
			}
			nodes = append(nodes, subnodes...)
		case *TextInputNode, *NumberInputNode:
			nodes = append(nodes, n)
		case *PreviewNode:
			nodes = append(nodes, n)
		default:
			return []Node{}, ErrFlattenUnknownNodeTypeFound
		}
	}
	return nodes, nil
}

func (g GroupNode) RelationList() (RelationList, error) {
	var ancestorStack []Node

	return g.getNodeRelation(0, ancestorStack)
}

func (g GroupNode) getNodeRelation(depth int, ancestorStack []Node) ([]Relation, error) {
	rels := make([]Relation, 0)

	ancestorStack = append(ancestorStack, &g)
	for _, n := range g.Nodes {
		switch n := (*n).(type) {
		case *GroupNode:
			subRels, err := n.getNodeRelation(depth+1, ancestorStack)
			if err != nil {
				return []Relation{}, err
			}
			rels = append(rels, subRels...)
		case *TextInputNode, *NumberInputNode, *PreviewNode:
			subRels := getNodeRelationFromInputNode(depth+1, ancestorStack, n)
			rels = append(rels, subRels...)
		default:
			return []Relation{}, ErrRelationListUnknownNodeTypeFound
		}
	}

	for i, v := range ancestorStack {
		rel := Relation{
			AncestorID:   v.ID(),
			DescendantID: g.ID(),
			Distance:     depth - i,
		}
		rels = append(rels, rel)
	}

	return rels, nil
}

func getNodeRelationFromInputNode(depth int, ancestorStack []Node, current Node) []Relation {
	rels := make([]Relation, 0)

	ancestorStack = append(ancestorStack, current)
	for i, v := range ancestorStack {
		rel := Relation{
			AncestorID:   v.ID(),
			DescendantID: current.ID(),
			Distance:     depth - i,
		}
		rels = append(rels, rel)
	}

	return rels
}
