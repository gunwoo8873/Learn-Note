package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		log.Fatalf("Not error : %v", err)
	}

	switch {
	case res.StatusCode == 600:
		fmt.Println("Unknown")
	case res.StatusCode == 500:
		fmt.Println("Server Error")
	case res.StatusCode == 400:
		fmt.Println("Client Error")
	case res.StatusCode == 300:
		fmt.Println("Redirection")
	case res.StatusCode == 200:
		fmt.Println("Success")
		goto exit
	case res.StatusCode == 100:
		fmt.Println("Informational")
	default:
		fmt.Println("Incorrect")
	}

exit:
	fmt.Println("Exit")
}
