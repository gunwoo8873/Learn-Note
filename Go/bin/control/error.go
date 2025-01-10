package control

import (
	"fmt"
	"net/http"
)

func makeCall(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error making GET request: %v", err)
	}

	return resp, nil
}

// Error case : Log save and continue, Error ignore, Stop and Panic, Global function to handle error
