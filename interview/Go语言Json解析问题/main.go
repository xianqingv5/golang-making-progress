package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

const jsonStream = `
    {
      "sounds/ambient/cave/cave1.ogg": {
        "hash": "fd6f6042f2a4ca8011e16c5199221c16db897f26",
        "size": 2338482
       }
     }
    {
      "sounds/ambient/cave/cave2.ogg": {
        "hash": "fd6f6042f2a4ca8011e16c5199221c16db897f27",
         "size": 1876295
        }
    }`

// Text text
type Text struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}

// Info info
type Info map[string]Text

func main() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var info Info
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for i, v := range info {
			fmt.Printf("%s : %s\n", i, v.Hash)
			fmt.Println(v.Size)
		}
	}
}
