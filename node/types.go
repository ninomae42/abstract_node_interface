package node

import "fmt"

type Type string

const (
	NodeTypeGroup Type = "NODE_GROUP"
	NodeTypeInput Type = "NODE_INPUT"
)

type GroupLevel string

const (
	GroupLevel1 GroupLevel = "GROUP_LEVEL_1"
	GroupLevel2 GroupLevel = "GROUP_LEVEL_2"
	GroupLevel3 GroupLevel = "GROUP_LEVEL_3"
	GroupLevel4 GroupLevel = "GROUP_LEVEL_4"
	GroupLevel5 GroupLevel = "GROUP_LEVEL_5"
)

func PrintNode(n Node) {
	switch n := n.(type) {
	case *GroupNode:
		fmt.Println("GroupNode")
		fmt.Printf("ID: %v, Label: %v, Level: %v, DirectionIsVertical: %v\n", n.ID(), n.Label, n.Level, n.DirectionIsVertical)
		for _, node := range n.Nodes {
			if node != nil {
				PrintNode(*node)
			}
		}
	case *TextInputNode:
		fmt.Println("TextInputNode")
		fmt.Printf("ID: %v, Label: %v, Required: %v, MaxLength: %v\n", n.ID(), n.Label, n.Required, n.MaxLength)
	case *NumberInputNode:
		fmt.Println("NumberInputNode")
		fmt.Printf("ID: %v, Label: %v, Required: %v, Unit: %v\n", n.ID(), n.Label, n.Required, n.Unit)
	default:
		fmt.Println("Unknown Node")
	}
}
