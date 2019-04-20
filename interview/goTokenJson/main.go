// Go 基于Token解析JSON
// https://blog.csdn.net/u011304970/article/details/72689121
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
		{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
	`
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := dec.Token() // 逗号和冒号被忽略
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}
