
//  tree.go
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

package tree

import (
	"fmt"
	"sync"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	root *node
	size int
	lock sync.Mutex
}

type node struct {
	Value interface{}
	Visited bool
	Adjacents []*Tree
}

// New returns a new, empty binary tree
func New() *Tree {
	t := &Tree{}
	t.lock = sync.Mutex{}
	return t
}	

func (t *Tree) Insert(value, parent interface{}){
	t.lock.Lock()
	defer t.lock.Unlock()
}

func (t *Tree) Remove(value interface{}){
	t.lock.Lock()
	defer t.lock.Unlock()
}

func (t *Tree) Size(){
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.size
}

/*
//NOTE - Changes required
func insert(t *Tree, v int) *Tree {
	if t == nil {
		//var adjacents []Tree  // an empty list
		//return &Tree{ v, adjacents}
		return &Tree{0, nil}

	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}
*/

func main() {

fmt.Println("Test main")
}