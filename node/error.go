package node

import "errors"

var (
	ErrBuildNodeTreeNodeOrParentChildRelationEmpty = errors.New("node: nodes or relations empty at BuildNodeTree")

	ErrBuildNodeTreeUnknownNodeTypeFound = errors.New("node: unknown node type found at BuildNodeTree")

	ErrBuildNodeTreeRootNodeNotFoundInNodes = errors.New("node: root node not found in nodes at BuildNodeTree")

	ErrNodeSpecifiedInRelationsNodeFoundInNodes = errors.New("node: node specified in relations not found in nodes at BuildNodeTree")

	ErrParentNodeSpecifiedInRelationsIsNotGroupNode = errors.New("node: parent node specified in relaions is not GroupNode at BuildNodeTree")

	ErrFlattenUnknownNodeTypeFound = errors.New("node: unknown node type found at Flatten")

	ErrRelationListUnknownNodeTypeFound = errors.New("node: unknown node type found at RelationList")
)
