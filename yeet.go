package main

import (
	"fmt"
	"os"
)

func main() {
	arg := "yeet"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	for {
		fmt.Println(arg)
	}
}
