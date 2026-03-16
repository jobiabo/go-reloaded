package main

import (
	"fmt"
	"helper/helper"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("commandline insufficient")
		return
	}

	inputFile := args[1]
	outputFile := args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}

	sampleIn := string(data)

	
	sampleIn = helper.HexBin(sampleIn)
	sampleIn = helper.CapAlphaMod(sampleIn)
	sampleIn = helper.Vowels(sampleIn)
	sampleIn = helper.PunctAAte(sampleIn)
	

	err = os.WriteFile(outputFile, []byte(sampleIn+"\n"), 0644)
	if err != nil {
		fmt.Println("error writing file", err)
		return
	}

	fmt.Println("synced successfully")

}
