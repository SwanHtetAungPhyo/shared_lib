package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer func(client *gosseract.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)
	err := client.SetLanguage("eng")
	if err != nil {
		return
	}
	err = client.SetImage("./IMG_4404.jpg")
	if err != nil {
		return
	}
	text, err := client.Text()
	if err != nil {
		return
	}
	fmt.Println(text)

}
