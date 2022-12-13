package main

import (
	"fmt"
	"github.com/ErfanMomeniii/del"
)

func main() {
	e := del.NewEvent("first")

	l1 := del.NewListener(func() {
		fmt.Println("hello")
	})

	l2 := del.NewListener(func() {
		fmt.Println("good By")
	})

	e.AddListener(l1, l2)

	e.Dispatch()
}
