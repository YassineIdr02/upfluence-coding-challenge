package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func ReadSSEPosts(postCh chan<- Post, duration time.Duration) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://stream.upfluence.co/stream", nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error connecting to SSE stream:", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	timeout := time.After(duration)

	for {
		select {
		case <-timeout:
			close(postCh)
			return
		default:
			if scanner.Scan() {
				line := scanner.Text()
				if len(line) < 6 || line[:5] != "data:" {
					continue // skip non-data lines
				}
				data := line[5:] // remove "data:" prefix

				// Parse the dynamic event
				var raw map[string]json.RawMessage
				if err := json.Unmarshal([]byte(data), &raw); err != nil {
					log.Println("Error unmarshaling event:", err)
					continue
				}

				for _, v := range raw {
					var post Post
					if err := json.Unmarshal(v, &post); err != nil {
						log.Println("Error decoding post:", err)
						continue
					}
					// Send the post to the channel
					postCh <- post
				}
			} else if err := scanner.Err(); err != nil {
				log.Println("Scanner error:", err)
			}
		}
	}
}