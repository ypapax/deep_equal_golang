package main

import (
	"log"
	"reflect"
)

func main() {
	deepEqualPrint(1, 1) // true
	deepEqualPrint(1, 2) // false
	var a = 3
	var b = 3
	deepEqualPrint(&a, &b) // true
	var c = 4
	var d = 5
	deepEqualPrint(&c, &d) //false
}

func deepEqualPrint(a, b any) {
	log.Println(a, b, reflect.DeepEqual(a, b))
}
