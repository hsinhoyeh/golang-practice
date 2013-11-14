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
    }
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

// PrintTree prints the nodes of binary tree with proper indentation and level.
func PrintTree(r *Node) {
    m := make(map[Value]int)
    InOrder(r, m)

    //do level-order traversal
    var treeBuffer *list.List
    var treeBNBuffer *list.List

    treeBNBuffer  =list.New()
    treeBNBuffer.PushBack(r)
    var preindent int;
    for true{
        //We do double buffer here: all the child pointers in the backend buffer (named with BN)
        //and then swap it when we go to a new iteration.
        treeBuffer = treeBNBuffer
        treeBNBuffer =list.New()
        preindent =0
        for b := treeBuffer.Front(); b != nil; b =b.Next(){
            bt := b.Value.(*Node)
            if bt == nil {
                continue
            }
            //The position in the map is the number of white space we want to print
            indent := m[bt.val]
            for i:=0; i < indent - preindent; i++ {
                fmt.Print(",")
            }
            fmt.Print(fmt.Sprintf("%c",bt.val))
            preindent = indent
            treeBNBuffer.PushBack(bt.left)
            treeBNBuffer.PushBack(bt.right)
        }
        fmt.Println()
        //break when there is no more element
        if treeBNBuffer.Len() ==0 {
            break;
        }
    }
}


//Build an map where key is the node.val and value is the position number
func InOrder(r *Node, m map[Value] int) {
    if r.left != nil{
        InOrder(r.left, m)
    }
    m[r.val]=len(m)+1
    if r.right != nil{
        InOrder(r.right, m)
    }
}

func main() {
    PrintTree(NewBinaryTree([]Value("ABCDEFGHI"), []Value("ABCDEFGHI")))
    PrintTree(NewBinaryTree([]Value("FBADCEGIH"), []Value("ABCDEFGHI")))
}
