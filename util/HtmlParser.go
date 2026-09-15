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
	if index+1 < inp_length {
		if (*inp)[index+1:index+4] == "!--" {
			isComment = true
		}
		else if (*inp)[index+1] == "!" {
			isDocDescriptor = true
		}
		else if (*inp)[index+1] == "/" {
			isClosing = true
		}
	}
}