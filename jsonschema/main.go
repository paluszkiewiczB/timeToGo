package main

import (
	"bytes"
	"encoding/json"
	"log"
)

const input = `
{
    "myObject": {
        "myObject": {
            "myString": "someValue"
        }
    }
}
`

func main() {
	in := &bytes.Buffer{}
	in.WriteString(input)

	out := &SampleJson{}
	err := json.NewDecoder(in).Decode(out)
	if err != nil {
		panic(err)
	}

	log.Printf("out: %+v", out)
}
