======================================================
/*

Write a function that constructs a binary tree from given in order and pre order traversals. Write a function that will print the binary tree with proper indentation and level.

For example:

IN:  ABCDEFGHI
PRE: ABCDEFGHI

A
    B
        C
            D
                E
                    F
                        G
                            H
                                I
PRE: FBADCEGIH
IN:  ABCDEFGHI
                    F
    B                   G
A           D                   I
        C       E           H
*/
package main

import ()

// Value defines the value type of a tree node.
type Value byte

// Node defines the binary tree node. Value is a byte.
type Node struct {
left  *Node
val   Value
right *Node
}

// NewBinaryTree constructs a binary tree from the given pre order
// and in order traversals. You can assume all inputs are valid. Empty []Value
// is considered valid valid.
func NewBinaryTree(pre, in []Value) *Node {
        return nil
}

// PrintTree prints the nodes of binary tree with proper indentation and level.
func PrintTree(r *Node) {
}

func main() {
PrintTree(NewBinaryTree([]Value("ABCDEFGHI"), []Value("ABCDEFGHI")))
PrintTree(NewBinaryTree([]Value("FBADCEGIH"), []Value("ABCDEFGHI")))
*/
