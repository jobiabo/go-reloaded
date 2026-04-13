package main

import (
	"fmt"
	"helper/helper"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("err: usage <go run .> <input.txt> <result.txt>")
	}

	inputRoute := args[1]
	outputRoute := args[2]

	read, err := os.ReadFile(inputRoute)
	if err != nil {
		fmt.Println(err)
	}
	// a := "i arrived safely"

	// read = []byte(a)

	stringForm := string(read)

	stringForm = helper.ModAphaNum(stringForm)

	err = os.WriteFile(outputRoute, []byte(stringForm), 0644)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Base conversion & Text transformation done succesfully")
	fmt.Println("Check file <result.txt> or use commmand <cat result.txt>")
}
