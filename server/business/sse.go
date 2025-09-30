package business

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"upfluence-coding-challenge/server/constants"
	"upfluence-coding-challenge/server/models"
)

// ReadSSEPosts fetches posts from the SSE stream sequentially for a given duration.
func ReadSSEPosts(duration time.Duration) ([]models.Post, error) {
	client := &http.Client{Timeout: duration + 5*time.Second}
	req, err := http.NewRequest("GET", constants.RouteStream, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "text/event-stream")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	start := time.Now()
	var posts []models.Post
	var eventData strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { // End of event
			if eventData.Len() > 0 {
				var raw map[string]json.RawMessage
				if err := json.Unmarshal([]byte(eventData.String()), &raw); err == nil {
					for _, v := range raw {
						var post models.Post
						if err := json.Unmarshal(v, &post); err != nil {
							log.Println("Error unmarshaling post:", err)
							continue
						}
						posts = append(posts, post)
					}
				}
				eventData.Reset()
			}
			continue
		}

		if strings.HasPrefix(line, "data:") {
			eventData.WriteString(line[5:])
		}

		if time.Since(start) >= duration {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Scanner error:", err)
	}

	return posts, nil
}
