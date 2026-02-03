package go_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed gocli.txt
var gocli string

func TestString(t *testing.T) {
	fmt.Println(gocli)
}

//go:embed favicon.svg
var favicon []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("favicon_new.svg", favicon, fs.ModePerm)
	if err != nil {
		t.Error(err)
	}
}

//go:embed files/*.txt
var files embed.FS

func TestEmbedFS(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}
