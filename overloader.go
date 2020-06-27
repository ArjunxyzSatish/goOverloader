package main

import (
//	"fmt"
	"math/rand"
	"os"
	"strconv"
)


func main() {

	for true{
		var filename string
		filename = strconv.Itoa(rand.Int())
//		fmt.Println("Creating file ", filename)
		os.Create(filename)
	}
}

