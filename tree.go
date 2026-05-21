package tree

import "fmt"


type TreeNode struct{
	Text string
	Children []TreeNode
}

func renderTree(items []TreeNode, prefix string, parentIsLast, isRoot bool) string {
	var line, result, childPrefix string
	result = ""
	for i, node := range items {
		isLastItem := (i == len(items) - 1)
	    if isRoot{
			line = fmt.Sprintf("%s\n", node.Text)
			childPrefix = ""
		} else {
			switch isLastItem{
			case true:
				line = fmt.Sprintf("%s└── %s\n", prefix, node.Text)
				childPrefix = fmt.Sprintf("%s    ", prefix)
			case false:
				line = fmt.Sprintf("%s├── %s\n", prefix, node.Text)
				childPrefix = fmt.Sprintf("%s│   ", prefix)
			}
		}
		result += line
		if len(node.Children) > 0 {
			result += renderTree(node.Children, childPrefix, isLastItem, false)
		}
		if isLastItem && isRoot && len(node.Children) == 0 {
			result += fmt.Sprintf("%s\n", prefix)
		}
	}
	return result
}


func Tree(items []TreeNode) string {
	return renderTree(items, "", true, true)
}
