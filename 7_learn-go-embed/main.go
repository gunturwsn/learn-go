package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed thunder.png
var logos []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	/* run 'go build' in the 7_learn-go-embed */
	fmt.Println(version)

	err := ioutil.WriteFile("thunder_new.png", logos, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}
