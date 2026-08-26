package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Bar string
	Baz string
}

type FooXML struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
}

func main() {
	f := Foo{"Joe junior", "Hello Shabado"}
	b, _ := json.Marshal(f)
	fmt.Println(string(b))
	json.Unmarshal(b, &f)
}
