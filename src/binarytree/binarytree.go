
//  binarytree.go
//
//  Created by Dinesh Appavoo
//	Copyright (c) 2015 ICRL
//
//	Permission is hereby granted, free of charge, to any person obtaining a copy
//	of this software and associated documentation files (the "Software"), to deal
//	in the Software without restriction, including without limitation the rights
//	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//	copies of the Software, and to permit persons to whom the Software is
//	furnished to do so, subject to the following conditions:
//
//	The above copyright notice and this permission notice shall be included in
//	all copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//	THE SOFTWARE.
//
// Go's concurrency primitives make it easy to
// express concurrent concepts, such as
// this binary tree comparison.
//
// Trees may be of different shapes,
// but have the same contents. For example:
//
//        4               6
//      2   6          4     7
//     1 3 5 7       2   5
//                  1 3
//
// This program compares a pair of trees by
// walking each in its own goroutine,
// sending their contents through a channel
// to a third goroutine that compares them.

package binarytree

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	root *node
	size int
	lock  *sync.Mutex
}

//Node struct for binary tree
type node struct {
	Left   *node
	Value  interface{}
	Right  *node
	Parent *node
	Visited bool
}

/

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func NewTree(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

// New returns a new, empty binary tree
func New() *Tree {
	t := &Tree{}
	t.lock = &sync.Mutex{}
	return t
}

//NewMinimalHeightBST gives the minimal height tree from the given sorted array
func NewMinimalHeightBST(arr []int, low int, high int) *Tree {
	if high < low {
		return nil
	}
	mid := (low + high) / 2
	t1 := NewTree()
	t1.Value = arr[mid]
	t1.Left = NewMinimalHeightBST(arr, low, mid-1)
	t1.Right = NewMinimalHeightBST(arr, mid+1, high)
	return t1
}

//Function to insert a new node into the tree
func (t *Tree) insert(t *Tree, v int) *Tree {
	if t == nil {
		t1 := &Tree{}
		t1.Parent = t
		return t1
	}
	if t.Left == nil && t.Right == nil {
		if v < t.Value {
			t.Left = &Tree{nil, v, nil, t, false}
		} else {
			t.Right = &Tree{nil, v, nil, t, false}
		}
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

//Function to remove a node
func (t *Tree) Remove() interface{} {

}

//Function to get a the value
func (t *Tree) Get() interface{} {

}

//Traverse prints the tree
func InOrderTraverse(t *Tree) {
	if t == nil {
		return
	}
	InOrderTraverse(t.Left)
	fmt.Print(t.Value, " ")
	InOrderTraverse(t.Right)
}

//Height gives the height of the BST
func Height(t *Tree) int {
	if t == nil {
		return 0
	}
	height := math.Max(Height(t.Left), Height(t.Right)) + 1
	return int(height)
}



func main() {
	//t1 := New(100, 1)
	//fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
	//fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
	//fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
	//fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")

	inArr := []int{4, 5, 7, 8, 9}
	t1 := NewMinimalHeightBST(inArr, 0, len(inArr)-1)
	InOrderTraverse(t1)
}
