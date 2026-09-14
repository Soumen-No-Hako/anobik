package parser

type HtmlNode struct{
	Tagname String
	TagId String
	TagClass String
	Txt String
	Next *HtmlNode
}

type TagNode struct{
	Tagname String
	TagId String
	TagClass String
	Txt String
	Next *TagNode
}

type XmlNode struct{
	Tagname String
	TagId String
	TagClass String
	Txt String
	Next *XmlNode
}