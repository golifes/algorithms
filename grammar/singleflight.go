package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	// We need a Group to use singleflight.
	var requestGroup singleflight.Group

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		// This time we'll wrap the githubStatus() call with singleflight's Group.Do()
		// Do takes a key (more on this later) and a function that returns a interface{} and an error.
		v, err, shared := requestGroup.Do("github", func() (interface{}, error) {
			// githubStatus() returns string, error, which statifies interface{}, error, so we can return the result directly.
			return githubStatus()
		})
		// Do returns an interface{}, error, and a bool which indicates whether multiple calls to the function shared the same result.

		// Check the error, as before.
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// We know that v will be a string, so we'll use a type assertion.
		status := v.(string)

		// Update the log statement so we can see if the results were shared.
		log.Printf("/github handler requst: status %q, shared result %t", status, shared)

		fmt.Fprintf(w, "GitHub Status: %q", status)
	})

	http.ListenAndServe("127.0.0.1:8081", nil)
}

// githubStatus retrieves GitHub's API status
func githubStatus() (string, error) {
	// No changes made to this function other than removing the comments for brevity.
	log.Println("Making request to GitHub API")
	defer log.Println("Request to GitHub API Complete")

	time.Sleep(1 * time.Second)

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("github response: %s", resp.Status)
	}

	r := struct{ Status string }{}

	err = json.NewDecoder(resp.Body).Decode(&r)

	return r.Status, err
}
