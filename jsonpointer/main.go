package main

import (
	"encoding/json"
	"fmt"

	"github.com/mattn/go-jsonpointer"
)

func main() {
	data := `
{
	"foo": [1, true, 2]
}
`
	var obj interface{}
	json.Unmarshal([]byte(data), &obj)
	rv, _ := jsonpointer.Get(obj, "/foo/1")
	rv2, _ := jsonpointer.Get(obj, "/foo/3")

	fmt.Printf("%v\n%v\n", rv, rv2)
}
