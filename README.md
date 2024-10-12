# Twitter API Go Application

A simple Twitter API application built in Go, allowing users to post and delete tweets through HTTP endpoints.

## Table of Contents

- [Introduction](#introduction)
- [Setup Instructions](#setup-instructions)
  - [Create a Twitter Developer Account](#create-a-twitter-developer-account)
  - [Generate API Keys](#generate-api-keys)
  - [Run the Program](#run-the-program)
- [Program Details](#program-details)
  - [Posting a Tweet](#posting-a-tweet)
  - [Deleting a Tweet](#deleting-a-tweet)
- [Error Handling](#error-handling)
- [Conclusion](#conclusion)

## [Introduction](#introduction)
This application demonstrates how to interact with the Twitter API using Go. It allows users to post and delete tweets via HTTP endpoints. The assignment teaches authentication with the Twitter API, handling errors gracefully, and creating a simple web server in Go.

## [Setup Instructions](#setup-instructions)

### 1. [Create a Twitter Developer Account](#create-a-twitter-developer-account)
- Go to the [Twitter Developer Platform](https://developer.twitter.com/en/apply-for-access).
- Sign in with your Twitter account and follow the prompts to create a developer account.

### 2. [Generate API Keys](#generate-api-keys)
- Once your account is set up, go to the [Twitter Developer Portal](https://developer.twitter.com/en/portal/dashboard).
- Create a new project and an app within that project.
- Under the "Keys and tokens" tab, you will find your `API Key`, `API Key Secret`, `Access Token`, and `Access Token Secret`.

### 3. [Run the Program](#run-the-program)
1. Install Go from the [official site](https://golang.org/dl/).
2. Clone this repository or download the source code.
3. Replace the placeholders for Twitter API credentials in the code with your actual credentials.
4. Run the application using:
   ```bash
   go run main.go
   ```

## [Program Details](#program-details)

### [Posting a Tweet](#posting-a-tweet)
- The `/tweet` endpoint accepts a POST request with a JSON body containing a `message` field. Example request:
   ```bash
   curl -X POST http://localhost:8080/tweet -H "Content-Type: application/json" -d '{"message": "Hello from Twitter API via localhost!"}'
   ```
- Successful response:
   ```json
   {
       "id": 1234567890,
       "text": "Hello from Twitter API via localhost!"
   }
   ```

### [Deleting a Tweet](#deleting-a-tweet)
- The `/delete` endpoint deletes the most recent tweet from the user's timeline. Example request:
   ```bash
   curl -X POST http://localhost:8080/delete
   ```
- Successful response:
   ```json
   {
       "id": 1234567890,
       "text": "Hello from Twitter API via localhost!"
   }
   ```

## [Error Handling](#error-handling)
- The application handles various Twitter API errors and sends JSON responses with meaningful error messages.
- Errors include invalid credentials, rate limits, and invalid tweet IDs. Each error is returned with a specific status code and message.

## [Conclusion](#conclusion)
This application serves as a basic introduction to working with the Twitter API using Go, demonstrating how to post and delete tweets while handling errors effectively.
