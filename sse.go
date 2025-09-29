package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// ReadSSEPosts reads posts from a Server-Sent Events (SSE) stream and sends them
// to the postCh channel. The function automatically stops after the specified duration.
func ReadSSEPosts(postCh chan<- Post, duration time.Duration) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://stream.upfluence.co/stream", nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error connecting to SSE stream:", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	timeout := time.After(duration)

	var eventData string

	for {
		select {
		case <-timeout:
			close(postCh)
			return
		default:
			if scanner.Scan() {
				line := scanner.Text()

				// Blank line signals end of an event
				if line == "" {
					if eventData != "" {
						processEvent(eventData, postCh)
						eventData = ""
					}
					continue
				}

				// Only lines starting with "data:" are part of the event
				if len(line) >= 5 && line[:5] == "data:" {
					eventData += line[5:]
				}
			} else if err := scanner.Err(); err != nil {
				log.Println("Scanner error:", err)
			}
		}
	}
}

// processEvent parses the JSON event and sends Post structs to the channel
func processEvent(eventData string, postCh chan<- Post) {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal([]byte(eventData), &raw); err != nil {
		log.Println("Error unmarshaling event:", err)
		return
	}

	for _, v := range raw {
		var post Post
		if err := json.Unmarshal(v, &post); err != nil {
			log.Println("Error decoding post:", err)
			continue
		}
		postCh <- post
	}
}