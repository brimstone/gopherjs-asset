package main

import (
	"flag"
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

var clientDir = "client"
var assetDir = "assets"

var assets http.FileSystem

func gopherjsMain() {

	options := &gbuild.Options{
		CreateMapFile: true,
		Verbose:       false,
		Minify:        true,
	}
	s := gbuild.NewSession(options)
	currentDirectory, err := os.Getwd()
	sourceDirectory := filepath.Join(currentDirectory, clientDir)
	assetsDirectory := filepath.Join(currentDirectory, assetDir)

	err = s.BuildDir(sourceDirectory,
		sourceDirectory,
		filepath.Join(assetsDirectory, "client.js"))

	fmt.Println(filepath.Join(assetsDirectory, "client.js"))

	if err != nil {
		log.Fatalln("error building directory", err)
	}
}

func assetsMain() {

	fmt.Println("Slurping up assets")
	err := vfsgen.Generate(assets, vfsgen.Options{
		BuildTags: "!dev",
	})
	if err != nil {
		log.Fatalln("Error running vfsgen:", err)
	}
}
func main() {

	flag.StringVar(&clientDir, "client", "client", "Directory for client package")
	flag.StringVar(&assetDir, "assets", "assets", "Directory for assets")
	flag.Parse()

	assets = union.New(map[string]http.FileSystem{
		"/" + assetDir: gopherjs_http.NewFS(http.Dir(assetDir)),
	})

	gopherjsMain()
	assetsMain()
}
