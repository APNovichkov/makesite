package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Post struct {
	Text string
}

func main() {
	startTime := time.Now()

	// Get flag value
	dirName := flag.String("dirName", "./posts", "Path to where post files are located")
	outDir := flag.String("outDir", "./html-templates", "Path to output html templates")
	flag.Parse()

	postPaths := getFilesInDirV2(*dirName)

	for postPath, size := range postPaths {
		fmt.Printf("Generating HTML file for post: %v with contents of size: %v bytes\n", filepath.Base(postPath), size)

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
		
		// Execute template
		err := t.Execute(f, post)
		if err != nil {
			panic(err)
		}
	
		f.Close()
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("Success! Generated %v html pages (%3.2f kb total) in %s.", len(postPaths), getSizeOfDir(*outDir), elapsedTime)
}

func getSizeOfDir(dirName string) float32 {
	var totalBytes int64 = 0
	err := filepath.Walk(dirName, 
		func(path string, info os.FileInfo, err error) error {
			if err != nil{
				panic(err)
			}

			if(!info.IsDir()){
				totalBytes += info.Size()
			}

			return nil
		})
	if err != nil{
		panic(err)
	}

	return convertBytesToKilobytes(totalBytes)
}

func convertBytesToKilobytes(bytes int64) float32{
	return float32(bytes) / 1000
}

func fileToString(filePath string) string {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func getFilesInDirV2(dirName string) map[string]int64 {
	// var outputPaths []string
	var outputPaths = make(map[string]int64)

	err := filepath.Walk(dirName, 
		func(path string, info os.FileInfo, err error) error{
			if err != nil{
				panic(err)
			}
			if(!info.IsDir()){
				outputPaths[path] = info.Size()
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



