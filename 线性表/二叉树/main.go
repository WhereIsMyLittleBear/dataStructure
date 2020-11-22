package main

import "fmt"

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

// TreeList 二叉树表
type TreeList struct {
	Tree *Tree
}

// Insert 插入节点
func (t *TreeList) Insert(data int) {
	if t.Tree == nil { //如果是空树，就插入到根节点
		t.Tree = &Tree{Data: data}
		return
	}
	reeNode := t.Tree
	for reeNode != nil {
		if data < reeNode.Data {
			if reeNode.Left == nil {
				reeNode.Left = &Tree{Data: data}
				return
			}
			reeNode = reeNode.Left
		} else if data > reeNode.Data {
			if reeNode.Right == nil {
				reeNode.Right = &Tree{Data: data}
				return
			}
			reeNode = reeNode.Right
		}
	}

}

// Pre 前序遍历
func (t *TreeList) Pre(reeNode *Tree) {
	if reeNode == nil {
		return
	}
	t.Pre(reeNode.Left)
	t.Pre(reeNode.Right)
}

// Mid 中序遍历
func (t *TreeList) Mid(reeNode *Tree) {
	if reeNode == nil {
		return
	}

	t.Mid(reeNode.Left)
	fmt.Printf("%#v ", reeNode.Data)
	t.Mid(reeNode.Right)
}

// After 后序遍历
func (t *TreeList) After(reeNode *Tree) {
	if reeNode == nil {
		return
	}
	t.After(reeNode.Left)
	t.After(reeNode.Right)
	fmt.Printf("%#v ", reeNode.Data)
}

//查找数据节点
func (t *TreeList) Select(data int) *Tree {
	if t.Tree == nil {
		return nil
	}
	reeNode := t.Tree
	for reeNode != nil {
		if data < reeNode.Data {
			reeNode = reeNode.Left
		} else if data > reeNode.Data {
			reeNode = reeNode.Right
		} else {
			return reeNode
		}
	}
	return nil
}

//删除数据节点
func (t *TreeList) Delete(data int) {
	if t.Tree == nil {
		return
	}

	reeNode := t.Tree
	var treeParent *Tree
	for reeNode != nil && data != reeNode.Data {
		treeParent = reeNode
		if data < reeNode.Data {
			reeNode = reeNode.Left
		} else {
			reeNode = reeNode.Right
		}
	}

	if reeNode == nil {
		return
	}

	if reeNode.Left != nil && reeNode.Right != nil {
		lastLeft := reeNode.Right //右子树的最小节点
		lastLeftParent := reeNode
		for lastLeft != nil && lastLeft.Left != nil {
			lastLeftParent = lastLeft
			lastLeft = lastLeft.Left
		}
		reeNode.Data = lastLeft.Data
		reeNode = lastLeft
		treeParent = lastLeftParent
	}

	var child *Tree
	if reeNode.Left != nil {
		child = reeNode.Left
	} else if reeNode.Right != nil {
		child = reeNode.Right
	}

	if treeParent == nil {
		t.Tree = child //删除的是根节点
	} else if treeParent.Left == reeNode {
		treeParent.Left = child
	} else if treeParent.Right == reeNode {
		treeParent.Right = child
	}
}

func main() {
	bsTree := &TreeList{}
	bsTree.Insert(33)
	fmt.Println(bsTree.Tree)
	bsTree.Mid(bsTree.Tree)

}
