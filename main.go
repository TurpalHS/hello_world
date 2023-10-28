package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	url := "http://api.openweathermap.org/data/2.5/find?q=yakutsk,RU&temp=like&humidity&APPID=4f594c4a2a73a44aebdf63f35837b419"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Print(string(byteValue))
}
