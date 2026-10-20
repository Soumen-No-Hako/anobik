package main

import (
	"fmt"
	"anobik/util"
	"os"
	"log"
)

func main() {
	htmlByteData, err := os.ReadFile("configs/test/htmlTestData.html")
	if err != nil {
    log.Fatal(err)
	}
	htmlData := string(htmlByteData)
	Test_parse_data_1(&htmlData)
}
func Test_parse_data_1(htmlInput *string) {
	extracted_info := anobikUtil.ParseHtml(htmlInput)
	// TODO A lot more conditions needed to properly filter and organise. Right now the data is gibberish and breaking in script tags and inline tags.
	// TODO Need a logic to identify parent tag of an inline to process or ignore tag and only process data
	fmt.Printf("Length of input = %d\n", len(*htmlInput))
	// fmt.Printf("Input value:\n %s\n", *htmlInput)
	fmt.Printf("Length of output = %d\n", len(extracted_info))
	fmt.Printf("Output : \n %s",extracted_info)
}