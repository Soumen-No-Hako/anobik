package anobikUtil

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

// Elements that break content flow into discrete structural blocks
var blockTags = map[string]struct{}{
	"p": {}, "div": {}, "article": {}, "section": {}, "main": {},
	"h1": {}, "h2": {}, "h3": {}, "h4": {}, "h5": {}, "h6": {},
	"li": {}, "blockquote": {}, "pre": {}, "tr": {}, "br": {}, "hr": {},
}

func parseHtml(resp string) string{
    var extractedData = ""
	var n = len(resp)
	i := 0
	for i<n {
		if resp[i] == "<" {
			isTag, isBlockTag, isDocDescriptor, isClosing, Tagname, i = getTagDetails(*resp, i, n)
			
		}
		
	}
}

func getTagDetails(inp *string, index int, inp_length int) (bool, bool, bool, bool, string, int) {

	isTag, isBlockTag, isDocDescriptor, isClosing, isComment := false, false, false, false, false
	if index+3<inp_length{
	if (*inp)[index+1:index+4] == "!--" {
			isComment = true
			getComment
		}
	}
	if index+1 < inp_length {
		if (*inp)[index+1] == "!" {
			isDocDescriptor = true
		}
		else if (*inp)[index+1] == "/" {
			isClosing = true
			forwardIndexToEnd
		}
		else {
			tag, index := detectTag(inp, index, inp_length)
			if blockTags[tag]
		}
	}
}

func detectTag(inp *string, index int, inp_length int) (string, int) {
	i := index
	for i < inp_length {
		if(*inp)[i] == ">" || (*inp)[i]==" " {
			return detectedTag, i;
		}
		detectedTag += (*inp)[i]
		i++
	}
	return "",inp_length;
}

func forwardIndexToEnd(inp *string, index int, inp_length int) int {
	// Start at index. keep moving right until you get >
	for index < inp_length {
		if(*inp)[index] == ">" {
			return index
		}
		index++
	}
	return inp_length
}

func getContent(inp *string, index int, inp_length int) (string, int) {
	string data := ""
	for index < inp_length {
		if(*inp)[index] == "<" {
			// Check it's a tag or math expression. if former then stop else continue
			isTag
			return data, index
		}
		data = data + (*inp)[index]
		index++
	}
	return inp_length
}