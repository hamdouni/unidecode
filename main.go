package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unidecode/libdecode"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Drag and drop a file or give filepath in parameter!")
		waitForKey()
		os.Exit(1)
	} else if !isFile(os.Args[1]) {
		fmt.Println("File does not exist!")
		waitForKey()
		os.Exit(1)
	}

	file := os.Args[1]

	content, err := readFromFile(file)
	if err != nil {
		fmt.Println(err)
		waitForKey()
		os.Exit(1)
	}
	decoded := libdecode.Unidecode(string(content))
	writeToFile(decoded, file)
}
func waitForKey() {
	fmt.Print("\nPress Enter to continue...")
	var buf [1]byte
	os.Stdin.Read(buf[:])
}
func isFile(file string) bool {
	s, err := os.Stat(file)
	if err != nil || s.IsDir() {
		return false
	}
	return true
}
func writeToFile(data, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}
