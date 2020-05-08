package network

import (
	"log"
	"net/http"
	"time"
)

func Get(address string) (*http.Response, error) {
	start := time.Now()
	client := &http.Client{}
	result, err := client.Get(address)
	if err != nil {
		elapsed := time.Since(start).Seconds()
		log.Printf("error or timeout %s %fs", address, elapsed)
		return nil, err
	}
	return result, err
}
