package main

import (
	"io/ioutil"
)

// func main() {
// 	fileContents := readFile()
// 	fmt.Println(fileContents)

// 	fmt.Println("Writing File!")
// 	writeFile("Hello World this is a new File!!")
// }

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func writeFile(stringToWrite string) {
	bytesToWrite := []byte(stringToWrite)
	err := ioutil.WriteFile("new-file", bytesToWrite, 0644)
	if err != nil{
		panic(err)
	}
} 

func renderTemplate(){
	
}

