package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)

func main() {
	var text string
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("eng")
	client.SetImage("./img.png")
	text, _ = client.Text()
	fmt.Println(text)

	client.SetImage("./img_1.png")
	text, _ = client.Text()
	fmt.Println(text)

	client.SetImage("./img_2.png")
	text, _ = client.Text()
	fmt.Println(text)

	client.SetImage("./img_3.png")
	text, _ = client.Text()
	fmt.Println(text)
}
