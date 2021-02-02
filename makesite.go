package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Post struct {
	Text string
}

func main() {
	// Get flag value
	// filePath := flag.String("postPath", "first-post.txt", "Path to your post")
	// outputPath := flag.String("outputPath", "output-post.html", "Output path for html file")
	dirName := flag.String("dirName", "./posts", "Path to where post files are located")
	outDir := flag.String("outDir", "./html-templates", "Path to output html templates")
	flag.Parse()

	postPaths := getFilesInDirV2(*dirName)

	for _, postPath := range postPaths {
		// Get file contents
		fileContents := fileToString(postPath)
		post := Post{fileContents}
	
		// Init template
		t := template.Must(template.New("mvp_template.tmpl").ParseFiles("mvp_template.tmpl"))
	
		// Create output file if it does not exist
		outputPath := filepath.Join(*outDir, fmt.Sprintf("%v.html", strings.Split(filepath.Base(postPath), ".")[0]))
		
		f, createFileErr := os.Create(outputPath)
		if createFileErr != nil{
			panic(createFileErr)
		}
	
		err := t.Execute(f, post)
		if err != nil {
			panic(err)
		}
	
		f.Close()
	}
}

func fileToString(filePath string) string {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func getFilesInDirV2(dirName string) []string {
	var outputPaths []string

	err := filepath.Walk(dirName, 
		func(path string, info os.FileInfo, err error) error{
			if err != nil{
				panic(err)
			}
			if(!info.IsDir()){
				outputPaths = append(outputPaths, path)
			}
			return nil
		})
	if err != nil{
		panic(err)
	}

	return outputPaths
}

func getFilesInDirV1(dirName string) []string {
	var outputPaths []string
	
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		outputPath := dirName + "/" + f.Name()
		outputPaths = append(outputPaths, outputPath)
	}

	fmt.Printf("Output paths: %v", outputPaths)

	return outputPaths
}



