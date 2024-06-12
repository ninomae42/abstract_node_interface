package node

import (
	"github.com/ninomae42/node_interface/domain"
)

type (
	Relation struct {
		AncestorID   domain.ID
		DescendantID domain.ID
		Distance     int
	}
	ParentChildRelation struct {
		ParentID domain.ID
		ChildID  domain.ID
	}
	RelationList []Relation
)

func (rl RelationList) ExtractParentChildRelations() []ParentChildRelation {
	var filtered []ParentChildRelation

	for _, r := range rl {
		if r.Distance == 1 {
			filtered = append(filtered,
				ParentChildRelation{
					ParentID: r.AncestorID, ChildID: r.DescendantID,
				})
		}
	}
	return filtered
}

type nodeMap map[domain.ID]Node

func BuildNodeTree(nodes []Node, rels []ParentChildRelation, rootID domain.ID) (Node, error) {
	if len(nodes) == 0 || len(rels) == 0 {
		return nil, ErrBuildNodeTreeNodeOrParentChildRelationEmpty
	}

	nMap := make(nodeMap, len(nodes))
	for _, n := range nodes {
		switch n := n.(type) {
		case *GroupNode:
			gn := *n
			gn.Nodes = nil
			nMap[gn.ID()] = &gn
		case *TextInputNode:
			in := *n
			nMap[in.ID()] = &in
		case *NumberInputNode:
			in := *n
			nMap[in.ID()] = &in
		default:
			return nil, ErrBuildNodeTreeUnknownNodeTypeFound
		}
	}

	if _, ok := nMap[rootID]; !ok {
		return nil, ErrBuildNodeTreeRootNodeNotFoundInNodes
	}

	for _, rel := range rels {
		parent, ok1 := nMap[rel.ParentID]
		child, ok2 := nMap[rel.ChildID]
		if !ok1 || !ok2 {
			return nil, ErrNodeSpecifiedInRelationsNodeFoundInNodes
		}

		switch parent := parent.(type) {
		case *GroupNode:
			parent.Nodes = append(parent.Nodes, &child)
		default:
			return nil, ErrParentNodeSpecifiedInRelationsIsNotGroupNode
		}

	}

	return nMap[rootID], nil
}
