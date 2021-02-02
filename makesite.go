package main

import (
	"flag"
	"io/ioutil"
	"os"
	"text/template"
)

type Post struct {
	Text string
}

func main() {
	// Get flag value
	filePath := flag.String("postPath", "first-post.txt", "Path to your post")
	outputPath := flag.String("outputPath", "output-post.html", "Output path for html file")
	flag.Parse()

	// Get file contents
	fileContents := fileToString(*filePath)
	post := Post{fileContents}

	// Init template
	t := template.Must(template.New("mvp_template.tmpl").ParseFiles("mvp_template.tmpl"))

	// Create output file if it does not exist
	f, createFileErr := os.Create(*outputPath)
	if createFileErr != nil{
		panic(createFileErr)
	}

	err := t.Execute(f, post)
	if err != nil {
		panic(err)
	}

	f.Close()
}

func fileToString(filePath string) string {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

// func main() {
// 	todos := ToDo{"Andrey", []entry{{"hi", false}, {"hello", true}}}

// 	// Files are provided as a slice of strings.
// 	paths := []string{
// 		"template.tmpl",
// 	}

// 	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
// 	err := t.Execute(os.Stdout, todos)
// 	if err != nil {
// 	  panic(err)
// 	}
// }

