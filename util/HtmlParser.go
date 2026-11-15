package anobikUtil

import (
	"anobik/model/parser"
	"strings"
	"fmt"
	"log"
)

// var htmlEscaper = strings.NewReplacer(
	// `&`, "&amp;",
	// `'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	// `<`, "&lt;",
	// `>`, "&gt;",
	// `"`, "&#34;", // "&#34;" is shorter than "&quot;".
// )


/*
var htmlTags = map[string]struct{}{
	"a": {}, "abbr": {}, "address": {}, "area": {}, "article": {}, "aside": {},
	"audio": {}, "b": {}, "base": {}, "bdi": {}, "bdo": {}, "blockquote": {},
	"body": {}, "br": {}, "button": {}, "canvas": {}, "caption": {}, "cite": {},
	"code": {}, "col": {}, "colgroup": {}, "data": {}, "datalist": {}, "dd": {},
	"del": {}, "details": {}, "dfn": {}, "dialog": {}, "div": {}, "dl": {},
	"dt": {}, "em": {}, "embed": {}, "fieldset": {}, "figcaption": {},
	"figure": {}, "footer": {}, "form": {}, "h1": {}, "h2": {}, "h3": {},
	"h4": {}, "h5": {}, "h6": {}, "head": {}, "header": {}, "hr": {},
	"html": {}, "i": {}, "iframe": {}, "img": {}, "input": {}, "ins": {},
	"kbd": {}, "label": {}, "legend": {}, "li": {}, "link": {}, "main": {},
	"map": {}, "mark": {}, "meta": {}, "meter": {}, "nav": {}, "noscript": {},
	"object": {}, "ol": {}, "optgroup": {}, "option": {}, "output": {}, "p": {},
	"picture": {}, "pre": {}, "progress": {}, "q": {}, "rp": {}, "rt": {},
	"ruby": {}, "s": {}, "samp": {}, "script": {}, "section": {}, "select": {},
	"small": {}, "source": {}, "span": {}, "strong": {}, "style": {}, "sub": {},
	"summary": {}, "sup": {}, "table": {}, "tbody": {}, "td": {}, "template": {},
	"textarea": {}, "tfoot": {}, "th": {}, "thead": {}, "time": {}, "title": {},
	"tr": {}, "track": {}, "u": {}, "ul": {}, "var": {}, "video": {}, "wbr": {},
}

// Elements that inherently denote boilerplate or non-visible content
var ignoredContainers = map[string]struct{}{
	"script": {}, "style": {}, "noscript": {}, "header": {}, "footer": {}, "nav": {},
}
*/
// Flags for operations and navigation

var inTag = false
var inQuote = false
var DataExtractorActive = true
var ext_text_buff string; // Buffer for extracted text to insert in different logic sections

// Elements that break content flow into discrete structural blocks
var blockTags = map[string]struct{}{
	"a":{}, "span":{}, "b":{}, "i":{}, "time":{}, "button":{}, "address":{},
	"p": {}, "div": {}, "article": {}, "section": {}, "main": {},
	"h1": {}, "h2": {}, "h3": {}, "h4": {}, "h5": {}, "h6": {},
	"li": {}, "blockquote": {}, "pre": {}, "tr": {}, "br": {}, "hr": {},
}

var FormInputTags = map[string]struct{}{
	"button":{},
	"datalist":{},
	"fieldset":{},
	"form":{},
	"input":{},
	"label":{},
	"legend":{},
	"optgroup":{},
	"option":{},
	"output":{},
	"select":{},
	"textarea":{},
}

var VoidTags = map[string] struct{}{
	// "!DOCTYPE":{},
	"area":{},
	"base":{},
	"basefont":{},
	"bgsound":{},
	"br":{},
	"col":{},
	"command":{},
	"embed":{},
	"frame":{},
	"hr":{},
	"img":{},
	"input":{},
	"isindex":{},
	"keygen":{},
	"link":{},
	"meta":{},
	"nextid":{},
	"param":{},
	"plaintext":{},
	"source":{},
	"track":{},
	"wbr":{},
}

var SkipTags = map[string] struct{}{
	"style":{},
	"script":{},
}
// TODO: need logic to skip complete Script blocks

func ParseHtml(resp *string) string{
    var extractedData string
	var finalData string 
	var n = len(*resp)
	var domStack = parser.NewHtmlDOMStack()
	domStack.Top = &parser.HtmlNode{}
	domStack.Size = 0
	domStack.Top.ChildNodes = new([]parser.HtmlNode)
	// var domTree = parser.NewHtmlDomTreeNode()
	var domTreeRoot *parser.HtmlNode
	
	i := 0
	var Tagname string
	var node parser.HtmlNode
	
	for i<n {
		var currentRune = (*resp)[i]
		i++
		if currentRune == '<' && !inQuote {
			node = parser.HtmlNode{}
			// Need some validations for skip tags
			if !inTag {
				node.Txt = ext_text_buff
				ext_text_buff = ""
			}
			inTag = !inTag
		} else if currentRune == '>' && !inQuote {
			if inTag {
				// Split by ' '.
				// determine tag
				// push to stack if non-closing. else add ext_text_buff data to stackNode and pop
				Tagname, isClosing, isVoidTag := ExtractTagDetails(ext_text_buff)
				inTag = !inTag
				if !isVoidTag {
				if isClosing {
					log.Printf("popping current head {%s} and detected tagname={%s}", domStack.Top.Tagname, Tagname)
					domStack.Pop()
				} else {
					node.Tagname = Tagname
					node.Txt = ext_text_buff
					log.Printf("current head {%s}", domStack.Top.Tagname)
					if domStack.Size>0 {
					if domStack.Top.ChildNodes == nil {
						domStack.Top.ChildNodes = new([]parser.HtmlNode)
					}
					*domStack.Top.ChildNodes = append(*domStack.Top.ChildNodes, node)
					//log.Printf("[INFO] %s's child node is - %s", domStack.Top.Tagname,domStack.Top.ChildNodes)
					}
					domStack.Push(&node)
					if domTreeRoot == nil {
						domTreeRoot = domStack.Top
					}
				}
				}
			}
		} else {
			if currentRune == '"' {
				inQuote = !inQuote
			}
			ext_text_buff += string(currentRune)
		}
	}
	fmt.Printf("Dom Tree is - %s",parser.PrintDomTreeSchema(domTreeRoot,0))
	
	for i<n {
		if (*resp)[i] == '<' {
		/*
		* Go until '>'; make sure > is not in a quote;
		* Do split by '<singleblankSpace>. First cut is the tag.
		* Insert tag in stack.
		* start getting content again.
		* 
		* Again if < comes but this time its closing tag, then pop
		* prefer using a single linear iterator to split single space and quote detection
		*/
			if len(extractedData) > 0 {
			finalData += extractedData + "\n"
			}
			extractedData = ""
			Tagname, i = GetTagDetails(resp, i, n)
		}
		// if script skip entirely
		// absolutely go till </script>
		if _, ok := blockTags[Tagname]; ok {
		extractedData = extractedData + string((*resp)[i]) //possible conditions can be added
		}
		i++
	}
	if len(strings.TrimSpace(extractedData)) > 0 {
			finalData += extractedData + "\n"
	}
	return finalData
}

func ExtractTagDetails (inp string) (string, bool, bool) {
	tmpList := strings.Split(inp," ")
	_, ok := VoidTags[tmpList[0]]
	if strings.Contains(tmpList[0], "/") {
		return tmpList[0], true, ok
	}
	return tmpList[0], false, ok
}
func GetTagDetails(inp *string, index int, inp_length int) (string, int) {

	//isDocDescriptor, isClosing, isComment := false, false, false
	// if index+3<inp_length{
	// if (*inp)[index+1:index+4] == "!--" {
			// isComment = true
			// getComment
		// }
	// }
	var tag string
	if index+1 < inp_length {
		if (*inp)[index+1] == '/' {
			index = ForwardIndexToEnd(inp, index, inp_length)
		} else {
			tag, index = DetectTag(inp, index+1, inp_length)
			if (*inp)[index] == ' ' || (*inp)[index] == '>' {
				index = ForwardIndexToEnd(inp, index, inp_length)
			}
		}
	}
	return tag, index
}

func DetectTag(inp *string, index int, inp_length int) (string, int) {
	var detectedTag string
	i := index
	for i < inp_length {
		if(*inp)[i] == '>' || (*inp)[i]==' ' {
			return detectedTag, i;
		}
		detectedTag += string((*inp)[i])
		i++
	}
	return "",i;
}

func ForwardIndexToEnd(inp *string, index int, inp_length int) int {
	// Start at index. keep moving right until you get >
	for index < inp_length {
		if(*inp)[index] == '>' {
			return index+1
		}
		index++
	}
	return inp_length
}