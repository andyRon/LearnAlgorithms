package main

import (
	. "github.com/andyron/LearnAlgorithms/hello-algo/pkg"
	"testing"
)

func TestNewTreeNode(t *testing.T) {

}

func TestLevelOrder(t *testing.T) {
	root := SliceToTree([]any{1, 2, 3, 4, 5, 6, 7})
	PrintTree(root)
	// TODO
}

func TestInOrder(t *testing.T) {

}

func TestPostOrder(t *testing.T) {

}

func TestPreOrder(t *testing.T) {

}
