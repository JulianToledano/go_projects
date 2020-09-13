package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	// b := make([]byte, 2074)
	// n, err := f.Read(b)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("File size", n)
	// fmt.Println(string(b))

	// dat, _ := ioutil.ReadFile(os.Args[1])
	// fmt.Println(string(dat))
	io.Copy(os.Stdout, f)
}
