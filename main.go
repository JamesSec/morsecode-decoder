package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("test.jpg")
	text, _ := client.Text()
	fmt.Println(text)
}
