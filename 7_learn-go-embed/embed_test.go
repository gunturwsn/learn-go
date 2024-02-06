package learn_go_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed thunder.png
var logos []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("thunder_new.png", logos, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
