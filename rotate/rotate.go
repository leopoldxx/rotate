package main

import (
	"fmt"
	"os"

	"github.com/leopoldxx/rotate"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	raw := os.Args[1:]
	for i := range raw {
		fmt.Println(rotate.Rotate(raw[i], "."))
	}
}
