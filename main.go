package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

// Custom error response structure
type ErrorResponse struct {
    StatusCode int    `json:"status_code"`
    Message    string `json:"message"`
}

// Global Twitter client
var client *twitter.Client

// Set up Twitter API credentials and client
func initTwitterClient() {
    consumerKey := "your_consumer_key"
    consumerSecret := "your_consumer_secret"
    accessToken := "your_access_token"
    accessSecret := "your_access_secret"

    // Set up OAuth1
    config := oauth1.NewConfig(consumerKey, consumerSecret)
    token := oauth1.NewToken(accessToken, accessSecret)
    httpClient := config.Client(oauth1.NoContext, token)

    // Create a new Twitter client
    client = twitter.NewClient(httpClient)
}

// Helper function to send a JSON error response
func sendJSONErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    jsonResponse := ErrorResponse{
        StatusCode: statusCode,
        Message:    message,
    }
    json.NewEncoder(w).Encode(jsonResponse)
}

// Modified handleTwitterError to return JSON error responses
func handleTwitterError(w http.ResponseWriter, resp *http.Response, err error) {
    if err != nil {
        // Send a detailed JSON error response
        sendJSONErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
        return
    }

    // Handle specific status codes with JSON response
    switch resp.StatusCode {
    case 400:
        sendJSONErrorResponse(w, http.StatusBadRequest, "Bad request: Invalid parameters")
    case 401:
        sendJSONErrorResponse(w, http.StatusUnauthorized, "Unauthorized: Check your API keys and tokens")
    case 403:
        sendJSONErrorResponse(w, http.StatusForbidden, "Forbidden: You may not have permission to perform this action")
    case 404:
        sendJSONErrorResponse(w, http.StatusNotFound, "Not found: Tweet not found")
    case 429:
        sendJSONErrorResponse(w, http.StatusTooManyRequests, "Rate limit exceeded: Try again later")
    case 500:
        sendJSONErrorResponse(w, http.StatusInternalServerError, "Twitter internal server error: Try again later")
    default:
        sendJSONErrorResponse(w, resp.StatusCode, fmt.Sprintf("Unexpected error: %v", resp.Status))
    }
}

// Handler to post a tweet
func tweetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        sendJSONErrorResponse(w, http.StatusMethodNotAllowed, "Invalid request method")
        return
    }

    var tweetRequest struct {
        Message string `json:"message"`
    }

    // Decode JSON request body
    if err := json.NewDecoder(r.Body).Decode(&tweetRequest); err != nil {
        sendJSONErrorResponse(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    // Post a new tweet with the provided message
    tweet, resp, err := client.Statuses.Update(tweetRequest.Message, nil)

    // Check for errors and handle based on HTTP response
    if err != nil || resp.StatusCode != 200 {
        handleTwitterError(w, resp, err) // Pass 'w'
        return
    }

    // Return the posted tweet as JSON
    response := struct {
        ID   int64  `json:"id"`
        Text string `json:"text"`
    }{
        ID:   tweet.ID,
        Text: tweet.Text,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Handler to delete the last tweet
func deleteTweetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        sendJSONErrorResponse(w, http.StatusMethodNotAllowed, "Invalid request method")
        return
    }

    tweets, resp, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{Count: 1})

    // Check if tweets were retrieved correctly
    if err != nil || resp.StatusCode != 200 {
        handleTwitterError(w, resp, err) // Pass 'w'
        return
    }

    if len(tweets) == 0 {
        sendJSONErrorResponse(w, http.StatusNotFound, "No tweets to delete")
        return
    }

    tweetID := tweets[0].ID
    deletedTweet, resp, err := client.Statuses.Destroy(tweetID, nil)

    // Check if the tweet was deleted correctly
    if err != nil || resp.StatusCode != 200 {
        handleTwitterError(w, resp, err) // Pass 'w'
        return
    }

    // Return the deleted tweet information as JSON
    response := struct {
        ID   int64  `json:"id"`
        Text string `json:"text"`
    }{
        ID:   deletedTweet.ID,
        Text: deletedTweet.Text,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    // Initialize the Twitter client
    initTwitterClient()

    // Set up HTTP routes
    http.HandleFunc("/tweet", tweetHandler)
    http.HandleFunc("/delete", deleteTweetHandler)

    // Start server on localhost:8080
    fmt.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
