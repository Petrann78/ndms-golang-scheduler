package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// HTTP endpoint
	url := os.Getenv("TARGET_URL")
	if url == "" {
		panic("TARGET_URL not set")
	}

	intervalStr := os.Getenv("INTERVAL")
	if intervalStr == "" {
		intervalStr = "5" // default 5 minutes
	}
	interval, err := time.ParseDuration(intervalStr + "m")
	if err != nil {
		panic("invalid INTERVAL: " + err.Error())
	}

	var body []byte
	if bodyFile := os.Getenv("BODY_FILE"); bodyFile != "" {
		body, err = os.ReadFile(bodyFile)
		if err != nil {
			panic(err)
		}
	} else {
		panic("BODY_FILE not set (need a JSON file with request body)")
	}

	// Create a HTTP post request
	client := &http.Client{}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("request build error:", err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("request error:", err)
			continue
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
			fmt.Println("bad status:", res.Status)
			continue
		}
		respBody, _ := io.ReadAll(res.Body)
		fmt.Printf("success [%s]: %s\n", res.Status, string(respBody))

	}
}
