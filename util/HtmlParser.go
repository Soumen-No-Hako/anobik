package anobikUtil

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

// Elements that break content flow into discrete structural blocks
var blockTags = map[string]struct{}{
	"a":{}, "span":{}, "b":{}, "i":{}, "time":{}, "button":{}, "address":{},
	"p": {}, "div": {}, "article": {}, "section": {}, "main": {},
	"h1": {}, "h2": {}, "h3": {}, "h4": {}, "h5": {}, "h6": {},
	"li": {}, "blockquote": {}, "pre": {}, "tr": {}, "br": {}, "hr": {},
}

func ParseHtml(resp *string) string{
    var extractedData string
	var finalData string 
	var n = len(*resp)
	i := 0
	var Tagname string
	for i<n {
		if (*resp)[i] == '<' {
			if len(extractedData) > 0 {
			finalData += extractedData + "\n"
			}
			extractedData = ""
			Tagname, i = GetTagDetails(resp, i, n)
		}
		if _, ok := blockTags[Tagname]; ok {
		extractedData = extractedData + string((*resp)[i]) //possible conditions can be added
		}
		i++
	}
	if len(extractedData) > 0 {
			finalData += extractedData + "\n"
	}
	return finalData
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