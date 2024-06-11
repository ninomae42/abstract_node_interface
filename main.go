package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/ninomae42/node_interface/domain"
	"github.com/ninomae42/node_interface/fixture"
	"github.com/ninomae42/node_interface/node"
)

func main() {
	// node fixtures
	root := fixture.GroupNode(
		fixture.GroupNodeWithID(domain.ID("1")),
		fixture.GroupNodeWithLabel(domain.String("root")),
		fixture.GroupNodeWithLevel(node.GroupLevel1),
	)

	subs := []node.Node{
		fixture.GroupNode(
			fixture.GroupNodeWithID(domain.ID("1-1")),
			fixture.GroupNodeWithLabel(domain.String("sub1")),
			fixture.GroupNodeWithLevel(node.GroupLevel2),
			fixture.GroupNodeWithNodes(
				fixture.TextInputNode(
					fixture.TextInputNodeWithID(domain.ID("1-1-1")),
					fixture.TextInputNodeWithLabel(domain.String("input1")),
					fixture.TextInputNodeWithMaxLength(1),
				),
			),
		),
		fixture.GroupNode(
			fixture.GroupNodeWithID(domain.ID("1-2")),
			fixture.GroupNodeWithLabel(domain.String("sub2")),
			fixture.GroupNodeWithLevel(node.GroupLevel2),
		),
	}

	root.AddNode(subs)
	// root.AddNode([]node.Node{nil})  // invalid node type assign to test

	// get flattened nodes
	nodes, err := root.Flatten()
	if err != nil {
		log.Fatal(err)
	}

	// get node relations
	rels, err := root.RelationList()
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(rels, func(i, j int) bool {
		return rels[i].AncestorID < rels[j].AncestorID
	})

	fmt.Println("relations")
	for _, rel := range rels {
		fmt.Printf("AncestorID: %v, DescendantID: %v, Distance: %v\n", rel.AncestorID, rel.DescendantID, rel.Distance)
	}

	// extract parent child relations(only distance 1)
	pcRels := rels.ExtractParentChildRelations()
	fmt.Println("parent child relations")
	for _, rel := range pcRels {
		fmt.Printf("ParentID: %v, ChildID: %v\n", rel.ParentID, rel.ChildID)
	}

	// build node tree from nodes and parent child relations
	tree, err := node.BuildNodeTree(nodes, pcRels, root.NodeID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===== tree =====")
	node.PrintNode(tree)
}
