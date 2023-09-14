package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed src/logo.png
var logo []byte

func TestBinaryFile(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/*.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(a))

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	c, err := files.ReadFile("files/c.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(c))

}

//go:embed ./files/*.txt
var path embed.FS

func TestPatchMatcher(t *testing.T) {
	dirEntries, err := path.ReadDir("files")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			file, err := path.ReadFile("files/" + entry.Name())
			if err != nil {
				panic(err)
			}

			fmt.Println(string(file))
		}
	}
}
