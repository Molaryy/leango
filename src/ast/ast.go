package ast

import "leango/src/Token"

type ASTNode struct {
	Parent   *ASTNode
	Children *ASTNode
	Next     *ASTNode
	Token    Token.Token
}

func AddNewChildren(token Token.Token, ast *ASTNode) *ASTNode {
	newNode := &ASTNode{
		Token:    token,
		Parent:   nil,
		Children: nil,
		Next:     nil,
	}

	if ast == nil {
		return newNode
	}
	newNode.Parent = ast
	ast.Children = newNode
	return ast
}

func AddToExistingChildren(token Token.Token, ast *ASTNode) *ASTNode {
	newNode := &ASTNode{
		Token:    token,
		Parent:   nil,
		Children: nil,
		Next:     nil,
	}

	if ast == nil {
		return newNode
	}
	tmp := ast

	for ast.Next != nil {
		ast = ast.Next
	}
	newNode.Parent = tmp
	ast.Next = newNode
	ast = tmp
	return ast
}
