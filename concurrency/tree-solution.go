package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int){
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int){
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value 
		_walk(t.Right, ch)
	}
}

func treeMain() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Print(v)
	}
}