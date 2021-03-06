package main

import (
	"fmt"
)

func getSomething() interface{} {
	return "example of string"
}

func getType(any interface{}) string {
	switch any.(type) {
	case string:
		return "string"
	case int:
		return "untyped int"
	case int64:
		return "int64"
		// and many cases farther, if we want handle more types
	default:
		return "i dont know what is it"
	}
}
func main() {
	fmt.Println(getType(getSomething()))
}
