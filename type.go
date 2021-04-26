package uelevelclip

type Block struct {
	Class    string
	Option   map[string]string
	Children []Node
}

type Line string

type Node interface {
	encode(e *Encoder) error
}
