package parser

import "strings"

type HtmlNode struct{
	Tagname string
	Txt string
	Attributes map[string]string
	Next *HtmlNode
	ChildNodes *[]HtmlNode
}

func PrintDomTreeSchema(root *HtmlNode, depth int) string {
	result := root.Tagname
	for index, node := range *root.ChildNodes {
		result += strings.Repeat("    ", depth)+string(index)+" -> "+PrintDomTreeSchema(&node, depth+1)+string('\n')
	}
	return result
}

type HtmlDOMStack[T HtmlNode] struct{
	Top *HtmlNode
	Size int
}

type HtmlDomTreeNode struct{
	HtmlNode *HtmlNode
	ChildNodes []*HtmlNode
}

func NewHtmlDOMStack[T HtmlNode]() *HtmlDOMStack[T]{
	return &HtmlDOMStack[T]{}
}

func NewHtmlDomTreeNode[T HtmlNode]() *HtmlDomTreeNode{
	return &HtmlDomTreeNode{}
}

func (stk *HtmlDOMStack[T]) Push(node *HtmlNode) {
	if stk.Top != nil {
		node.Next = stk.Top
	}
	stk.Top = node;
	stk.Size = stk.Size + 1
}

func (stk *HtmlDOMStack[T]) Pop() {
	if stk == nil || stk.Top == nil {
		return
	}
	popped := stk.Top
	stk.Top = stk.Top.Next
	popped.Next = nil
	stk.Size--
}


type TagNode struct{
	Tagname string
	TagId string
	TagClass string
	Txt string
	Next *TagNode
}

type XmlNode struct{
	Tagname string
	TagId string
	TagClass string
	Txt string
	Next *XmlNode
}