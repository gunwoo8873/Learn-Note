package control

import (
	"fmt"
	"log"
	"net/http"
)

func Switch_Case() {
	rep, err := http.Get("localhost:8080")

	if err != nil {
		log.Fatalf("Could not Coonnect : %v", err)
	}

	switch {
	case rep.StatusCode >= 600:
		fmt.Printf("Unknown Error : %v\n", rep.StatusCode)
		goto exception
	case rep.StatusCode >= 400:
		fmt.Printf("Client Error : %v\n", rep.StatusCode)
		goto failure
	case rep.StatusCode >= 300:
		fmt.Printf("Redirection : %v\n", rep.StatusCode)
		goto failure
	case rep.StatusCode >= 200:
		fmt.Printf("Success : %v\n", rep.StatusCode)
		goto exit
	}

exception:
	panic("Exception")
failure:
	log.Fatalf("Failed to Connect: %v", err)
exit:
	fmt.Println("Connection Successful")
}

// nil : Null
