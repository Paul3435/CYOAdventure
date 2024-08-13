package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/Paul3435/CYOAdventure"
)

func main() {
	//Create flag for file to use
	fileLocation := flag.String("file", "stories.json", "Specify the file in which you want to run this program (Default: stories.json)")
	flag.Parse()

	//Read the file
	file := readFile(*fileLocation)

	story, err := cyoa.JsonReader(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}

func readFile(fileLocation string) *os.File {
	f, err := os.Open(fileLocation)
	if err != nil {
		panic(err)
	}

	return f
}
