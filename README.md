# tree
A simple, zero dependency go library for rendering tree structures on the terminal.

## Installation
```bash
go get github.com/jobotow/tree
```

## Usage
```go
package main
import (
    "fmt"
    "github.com/jobotow/tree"
)

func main(){
    data := []tree.TreeNode{
        {Text: "Root", Children: []tree.TreeNode{
            {Text: "Branch 1", Children:
            []tree.TreeNode{
                {Text: "Leaf 1"},
                {Text: "Leaf 2"},
            }},
            {Text: "Branch 2", Children:
            []tree.TreeNode{
                {Text: "Leaf 3"},
            }},
        }},
    }
	fmt.Println(tree.Tree(data))
}
```

## Output 
```bash
Root
├── Branch 1
│   ├── Leaf 1
│   └── Leaf 2
└── Branch 2
    └── Leaf 3
```

## Use Cases
- Filesystem trees
- Dependency graphs
- Any thing that needs output in hierachical form.

## LICENSE
Apache 2.0



