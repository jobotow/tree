package main
import (
	"fmt"

)

type TreeNode struct{
	Text string
	Children []TreeNode
}

func Tree(items []TreeNode, prefix string, parentIsLast, isRoot bool) string {
	//prefix = ""
	//parentIsLast =
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
				childPrefix = "    "
			case false:
				line = fmt.Sprintf("%s├── %s\n", prefix, node.Text)
				childPrefix = "│   "
			}
		}
		result += line
		if len(node.Children) > 0 {
			fmt.Printf("Child PREFIX: '%s'\n", childPrefix)
			result += Tree(node.Children, childPrefix, isLastItem, false)
		}
		if isLastItem && isRoot && len(node.Children) == 0 {
			result += fmt.Sprintf("%s\n", prefix)
		}
	}
	return result
}

func main(){
	myTree := []TreeNode{
		{Text: "Level 1", Children:
	[]TreeNode{
		{Text: "Level 2", Children:
	[]TreeNode{
		{Text: "Level 3", Children:
	[]TreeNode{
		{Text: "Level 4"}, // Children:
	//[]TreeNode{
	//	{Text: "Level 5"},
	//}},
	}},
	}},
	}},
	}
	fmt.Println(Tree(myTree, "", true, true))
	
}


/*
    def self.render_tree(items, prefix = "", parent_is_last=true, is_root=true)
      result = ""

      items.each_with_index do |node, i|
        is_last_item = (i == items.length - 1)

        line = if is_root
          "#{node[:text]}\n"
        else
          "#{HEITT::Color.colorize(prefix, :blue)}#{HEITT::Color.colorize((is_last_item ? '└── ' : '├── '), :blue)}#{node[:text]}\n"
        end

        child_prefix = if is_root
          ""
        else
          "#{HEITT::Color.colorize(prefix, :bold, :blue)}#{HEITT::Color.colorize((is_last_item ? "    " : "│   "), :bold, :blue)}"
        end
        result += line 
        result += render_tree(node[:children], child_prefix, is_last_item, false) 
		if node[:children].any?
      
        if is_last_item && !is_root and !node[:children].any?
          result += "#{HEITT::Color.colorize(prefix, :bold, :blue)}  \n"
        end
      end
      result
    end
/*/