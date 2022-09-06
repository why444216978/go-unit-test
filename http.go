package main

import (
	"fmt"
	"net/http"
	"time"
)

func Send() (err error) {
	req, err := http.NewRequest(http.MethodGet, "https://127.0.0.1:8080", nil)
	if err != nil {
		return
	}
	client := &http.Client{
		Timeout: time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status is %d", resp.StatusCode)
	}

	return
}
