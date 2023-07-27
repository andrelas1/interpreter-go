package ast

import "monkey/token"

type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node // this makes the Statement EXTEND a Node
  statementNode()
}

type Expression interface {
  Node // this makes the Expression EXTEND a Node
  expressionNode()
}

type Program struct {
  Statements []Statement // because this is a struct and not an interface, this Statements is actually a property of the Program
}

func (p *Program) TokenLiteral() string {
  if(len(p.Statements) > 0) {
    return p.Statements[0].TokenLiteral()
  } else {
    return ""
  }
}

type LetStatement struct {
  Token token.Token
  Name *Identifier
  Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
  return ls.Token.Literal
}

type Identifier struct {
  Token token.Token
  Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
  return i.Token.Literal
}
