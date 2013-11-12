/*
    this go program load two sequence as input and return a sequence of output.
    the inputs are a binary tree which is traversed in pre-order and in-order.
    the output is traversed in lever-order.
    this program is for demo only
*/

package main

import (
    "fmt"
    "container/list"
)

// Value defines the value type of a tree node.
type Value byte

// Node defines the binary tree node. Value is a byte.
type Node struct {
    left  *Node
    val   Value
    right *Node
}

func stringInSlice(t Value, seq []Value) int {
    for i, v := range seq {
        if t == v {
            return i
        }
    }
    //TODO: to avoid this case
    return -1
}

// NewBinaryTree constructs a binary tree from the given pre order
// and in order traversals. You can assume all inputs are valid. Empty []Value
// is considered valid valid.
func NewBinaryTree(pre, in []Value) *Node {

    if pre == nil || len(pre) ==0 {
        return nil
    } else {
        root := new(Node)
        root.val = pre[0]
        pos := stringInSlice(root.val, in)
        if pos < 0 {
            root.left = nil
            root.right = nil
        } else {
            root.left = NewBinaryTree(pre[1:pos+1], in[:pos])
            root.right = NewBinaryTree(pre[pos+1:], in[pos+1:])
        }
        return root
    }
}

func DepthMost(r *Node, depth int) int {
    if r == nil {
        return depth
    }
    if r.left == nil && r.right == nil {
        return depth
    } else {
        return MaxInt(DepthMost(r.left, depth +1), DepthMost(r.right, depth+1))
    }
}

func LeftDepthMost(r *Node, lcnt int) int {
    if r.left == nil {
        return lcnt
    } else {
        return LeftDepthMost(r.left, lcnt +1)
    }
}

func MaxInt(x, y int) int {
    if x > y {
        return x
    } else{
        return y
    }
}

func IndentBase(x, y int) int {
    return x 
}

func IndentNum(base, position int) int {
    //TODO: this array should be made in static
    arr := make([]int, position)
    arr[0] = 1
    for i:=1; i< position; i++ {
        arr[i] = arr[i-1] + IndentBase(base, i)
    }
    return arr[position-1]
}

// PrintTree prints the nodes of binary tree with proper indentation and level.
func PrintTree(r *Node) {
    var treeBNBuffer *list.List
    var treeBuffer *list.List
    var indentBNBuffer *list.List
    var indentBuffer *list.List

    maxDepth := DepthMost(r, 0)
    lDepth := LeftDepthMost(r, 0)
    depth := maxDepth
    treeBNBuffer  =list.New()
    indentBNBuffer =list.New()
    if lDepth == 0{
        //do not necessary shift the root
        indentBNBuffer.PushBack(0)
    } else{
        indentBNBuffer.PushBack(IndentNum(2, maxDepth))
    }
    treeBNBuffer.PushBack(r)

    for true {
        depth --
        //do swap and reallocate list
        treeBuffer = treeBNBuffer
        indentBuffer = indentBNBuffer
        treeBNBuffer  =list.New()
        indentBNBuffer =list.New()

        preindent := 0
        for b, i := treeBuffer.Front(), indentBuffer.Front(); b != nil; b, i =b.Next(), i.Next(){
            bt := b.Value.(*Node)
            it := i.Value.(int)
            if bt == nil {
                continue
            }
            
            for i:=0; i < it - preindent; i++ {
                fmt.Print(",")
            }
            fmt.Print(fmt.Sprintf("%q",bt.val))
            treeBNBuffer.PushBack(bt.left)
            indentBNBuffer.PushBack( it - IndentBase(2, depth))
            treeBNBuffer.PushBack(bt.right)
            indentBNBuffer.PushBack(it + IndentBase(2, depth))
            preindent = it
        }
        //newline
        fmt.Println()
        if treeBNBuffer.Len() ==0 {
            break;
        }
    }
}

func main() {
    PrintTree(NewBinaryTree([]Value("ABCDEFGHI"), []Value("ABCDEFGHI")))
    PrintTree(NewBinaryTree([]Value("FBADCEGIH"), []Value("ABCDEFGHI")))
}
