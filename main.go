package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shurcooL/go/gopherjs_http"
	"github.com/shurcooL/httpfs/union"
	"github.com/shurcooL/vfsgen"

	gbuild "github.com/gopherjs/gopherjs/build"
)

var assets = union.New(map[string]http.FileSystem{
	"/assets": gopherjs_http.NewFS(http.Dir("assets")),
})

func gopherjsMain() {

	options := &gbuild.Options{
		CreateMapFile: true,
		Verbose:       false,
		Minify:        true,
	}
	s := gbuild.NewSession(options)
	currentDirectory, err := os.Getwd()
	sourceDirectory := filepath.Join(currentDirectory, "client")
	assetsDirectory := filepath.Join(currentDirectory, "assets")

	err = s.BuildDir(sourceDirectory,
		sourceDirectory,
		filepath.Join(assetsDirectory, "client.js"))

	fmt.Println(filepath.Join(assetsDirectory, "client.js"))

	if err != nil {
		log.Fatalln(err)
	}
}

func assetsMain() {

	fmt.Println("Slurping up assets")
	err := vfsgen.Generate(assets, vfsgen.Options{
		BuildTags: "!dev",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	gopherjsMain()
	assetsMain()
}
