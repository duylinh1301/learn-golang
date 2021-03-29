package main

import (
	bootstrap "blog/bootstrap"
)

func main() {
	bootstrap.HandleRequests()
	// err := ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	// if err != nil {
	// 	fmt.Printf("Unable to write file: %v", err)
	// }
}
